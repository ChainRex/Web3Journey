package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (server *Server) ListenMessager() {
	for {
		msg := <-server.Message
		// 将msg发送给全部的在线user
		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}

// 广播消息
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	// 用户上线，将用户加入到OnlineMap中
	user := NewUser(conn, server)

	user.Online()

	// 监听用户是否活跃
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			msg := string(buf[:n-1])

			user.DoMessage(msg)
			// 用户针对客户端的任意消息，都应该重置定时器
			isLive <- true
		}
	}()
	for {
		// select没有default会一直阻塞
		select {
		// 当前用户活跃
		case <-isLive:
			// 不写代码，打破select阻塞状态，每个条件重新判断，定时器重新计时

		// 本质是一个channel
		case <-time.After(time.Second * 300):
			// 超时踢出
			user.SendMsg("you have been kick out\n")
			// 销毁资源
			close(user.C)
			conn.Close()
			return

		}

	}

}

// 启动服务器的接口
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go server.ListenMessager()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		// do handler
		go server.Handler(conn)
	}

}
