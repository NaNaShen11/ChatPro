package main

import (
	"chatPro/src/config"
	"chatPro/src/server/handle"
	"fmt"
	"net"
)


func talk(conn net.Conn) {
	defer conn.Close()
	handler :=handle.Handler{conn}
	handler.HandlerMain()
}

func main() {
	//获取配置
	host := config.Conf.Host
	//创建listener
	listener,err := net.Listen("tcp",host)
	if err!=nil{
		fmt.Println("server listen is err.")
		return
	}
	defer listener.Close()
	fmt.Println("server already")
	//创建监听
	for  {
		conn,err :=listener.Accept()
		if err!=nil{
			fmt.Println("create conn err .")
			return
		}
		//等待建立链接
		go talk(conn)
	}

}


