package main

import (
	"chatPro/src/config"
	"chatPro/src/server/accept"
	"fmt"
	"net"
)

type Client struct {
	C chan string
}

var onlineMap map[string]Client

var message = make(chan string)

func WriteToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg))
	}
}

func talk(conn net.Conn) {
	defer conn.Close()
	//创建管道
	clnt := Client{make(chan string)}
	onlineMap[conn.RemoteAddr().String()] = clnt
	accpet := accept.Accept{conn, clnt.C}
	accpet.Accpet()
	go WriteToClient(clnt, conn)
}

func Manager() {
	//初始化
	onlineMap = make(map[string]Client)
	for {
		msg := <-message
		for _, s := range onlineMap {
			s.C <- msg
		}
	}

}

func main() {
	//获取配置
	host := config.Conf.Host
	//创建listener
	listener, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("server listen is err.")
		return
	}
	defer listener.Close()
	fmt.Println("server already")

	go Manager()

	//创建监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("create conn err .")
			return
		}
		//等待建立链接
		go talk(conn)
	}

}
