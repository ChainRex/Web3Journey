package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前客户端的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client
}

// 处理server回应的消息，直接显示在标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. Public chat")
	fmt.Println("2. Private chat")
	fmt.Println("3. Rename")
	fmt.Println("0. Quit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>> Please input valid number")
		return false
	}
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>> Please input message, input 'exit' to quit")

	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送给server
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>> Please input message, input 'exit' to quit")
		fmt.Scanln(&chatMsg)
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {

	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>> Please input remote user name, input 'exit' to quit")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>> Please input message, input 'exit' to quit")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err:", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>> Please input message, input 'exit' to quit")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		remoteName = ""
		fmt.Println(">>>>> Please input remote user name, input 'exit' to quit")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) Rename() bool {
	fmt.Println(">>>>> Please input your new name:")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			// public chat
			fmt.Println(">>>>> Public chat")
			client.PublicChat()
			break
		case 2:
			// private chat
			fmt.Println(">>>>> Private chat")
			client.PrivateChat()
			break
		case 3:
			// rename
			fmt.Println(">>>>> Rename")
			client.Rename()
			break
		}
	}

}

var serverIp string
var serverPort int

func init() {
	// 绑定命令行参数
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set server ip address(default:127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "set server port(default:8888)")

}

func main() {
	// 解析命令行
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>> connect server failed")
		return
	}

	go client.DealResponse()

	fmt.Println(">>>>> connect server success")

	// 启动客户端业务
	client.Run()

}
