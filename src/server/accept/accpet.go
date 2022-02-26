package accept

import (
	"bufio"
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

type Accept struct {
	Conn net.Conn
	C    chan string
}

func (accept *Accept) Accpet() {
	//读取消息
	for {
		reader := bufio.NewReader(accept.Conn)
		//读取消息长度
		opType, data, err := util.Decode(reader)
		if err != nil {
			return
		}
		//解析数据
		var person protoc.Person
		err = proto.Unmarshal(data, &person)
		if err != nil {
			fmt.Println("unmarshal err:", err.Error())
			return
		}
		switch opType {
		case fmt.Sprint(LoginRequest):
			fmt.Println("opType:", opType)
			fmt.Println("name:", person.Name)
			fmt.Println("password:", person.Password)
			//登陆成功 存入客户端conn信息
		}
	}
}
