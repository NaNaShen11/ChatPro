package handler

import (
	"chatPro/src/proto/protoc"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func Login(msg []byte) {
	var person protoc.Person
	err := proto.Unmarshal(msg, &person)
	if err != nil {
		fmt.Println("unmarshal err:", err.Error())
		return
	}
	fmt.Println("name:", person.Name)
	fmt.Println("password:", person.Password)
}
