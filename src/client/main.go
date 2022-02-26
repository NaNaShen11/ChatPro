package main

import (
	"chatPro/src/proto/protoc"
	"chatPro/src/util"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

const (
	LoginRequest    = 1001
	LoginReply      = 2001
	SendChatRequest = 1002
	SendChatReply   = 2002
)

type Message struct {
	Data string
}

var (
	name     string
	password string
)

func main() {
	//建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("client dial err")
		return
	}
	fmt.Println("login input '1'")
	var input int
	fmt.Scanf("%d\n", &input)
	switch input {
	//login
	case 1:
		fmt.Println("input name:")
		fmt.Scanf("%s\n", &name)
		fmt.Println("login password:")
		fmt.Scanf("%s\n", &password)
		person := protoc.Person{Name: *proto.String(name), Password: *proto.String(password)}
		// 对数据进行序列化
		data, err := proto.Marshal(&person)
		if err != nil {
			fmt.Println("marshal err")
			return
		}
		//编码
		data, err = util.Encode(string(data), LoginRequest)
		//发送数据
		_, err = conn.Write(data)
		if err != nil {
			return
		}
		fmt.Println("senData,%s", data)
	}
}
