package handle

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

type Handler struct {
	Conn net.Conn
}

func (handler *Handler) HandlerMain()  {
	var buf [1024]byte
	//读取消息
	for  {
		//读取消息长度
		len,err :=handler.Conn.Read(buf[0:4])
		if err!=nil{
			return
		}
		datalen := binary.BigEndian.Uint32(buf[0:4])
		len,err = handler.Conn.Read(buf[0:datalen])
		if err!=nil{
			return
		}

		if len != int(datalen){
			return
		}
		//解析数据
		var cmd protoc.Cmd
		err =proto.Unmarshal(buf[:datalen],&cmd)
		if err!=nil{
			fmt.Println("unmarshal err")
			return
		}
		switch cmd.Id {
		case LoginRequest:
			var person protoc.Person
			err = proto.Unmarshal([]byte(cmd.Data),&person)
			if err !=nil{
				return
			}
			fmt.Println("name:",person.Name)
			fmt.Println("password:",person.Password)
			//登陆成功 存入客户端conn信息
		}
	}
}
