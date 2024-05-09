package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听当前user channel消息的协程
	go user.ListenMessage()
	return user
}

// 用户上线
func (user *User) Online() {
	// 用户上线，将用户加入到OnlineMap中

	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.server.BroadCast(user, "on line")
}

// 用户下线
func (user *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()
	// 广播当前用户下线消息
	user.server.BroadCast(user, "off line")
}

// 给指定用户发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户处理消息
func (user *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户
		user.server.mapLock.Lock()
		for _, onlineUser := range user.server.OnlineMap {
			onlineMsg := "[" + onlineUser.Addr + "]" + onlineUser.Name + ":" + "online\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("username has been used\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.Name = newName
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.SendMsg("rename success:" + newName + "\n")

		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|张三|消息内容

		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("msg format err, should be to|张三|msg content\n")
			return
		}

		remoteUser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("user not found\n")
			return

		}

		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("msg content is empty\n")
			return
		}
		remoteUser.SendMsg(user.Name + " send to you: " + content + "\n")

	} else {
		user.server.BroadCast(user, msg)
	}

}

// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}
