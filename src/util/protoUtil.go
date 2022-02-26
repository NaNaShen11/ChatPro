package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//编码
func Encode(message string) ([]byte, error) {
	//取消息长度(4字节)
	len :=int32(len(message))
	//建立缓存区
	buffer :=new(bytes.Buffer)
	//写入消息头
	err :=binary.Write(buffer,binary.BigEndian,len)
	if err!=nil{
		fmt.Println("binary write head err: ",err.Error())
		return nil, err
	}
	//写入消息体
	err = binary.Write(buffer,binary.BigEndian,[]byte(message))
	if err != nil{
		fmt.Println("binary write data err: ",err.Error())
		return nil,err
	}
	return buffer.Bytes(), nil
}

//解码
func UnEncode()  {
	
}
