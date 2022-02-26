package accept

import (
	"bufio"
	"chatPro/src/server/handler"
	"chatPro/src/util"
	"fmt"
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
		fmt.Println("opType:", opType)
		if err != nil {
			return
		}
		switch opType {
		case fmt.Sprint(LoginRequest):
			handler.Login(data)
		}
	}
}
