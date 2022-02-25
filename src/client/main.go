package main

import (
	"chatPro/src/proto/protoc"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

const (
	LoginRequest = 1001
	LoginReply = 2001
	SendChatRequest = 1002
	SendChatReply = 2002
)

type Message struct {
	Data string
}

var (
	name string
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
		cmd := protoc.Cmd{Id: LoginRequest, Data: string(data)}
		data, err = proto.Marshal(&cmd)
		//发送数据
		var dataLen uint32
		dataLen = uint32(len(data))
		var bytes [4]byte
		binary.BigEndian.PutUint32(bytes[0:4], dataLen)
		//发送消息长度
		n, err := conn.Write(bytes[:])
		if err != nil {
			return
		}
		//发送消息本身
		n, err = conn.Write(data)
		if err != nil {
			return
		}
		fmt.Println("senData,%v", n)
	}
}
