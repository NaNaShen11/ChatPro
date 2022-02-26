package util

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

//编码
func Encode(message string, opType uint16) ([]byte, error) {
	//取消息长度(4字节)
	len := int32(len(message))
	//建立缓存区
	buffer := new(bytes.Buffer)
	//写入消息长度(4字节)
	err := binary.Write(buffer, binary.BigEndian, len)
	if err != nil {
		fmt.Println("binary write head err: ", err.Error())
		return nil, err
	}
	//写入协议号（2字节）
	err = binary.Write(buffer, binary.BigEndian, opType)
	//写入消息体
	err = binary.Write(buffer, binary.BigEndian, []byte(message))
	if err != nil {
		fmt.Println("binary write data err: ", err.Error())
		return nil, err
	}
	return buffer.Bytes(), nil
}

//解码
func Decode(reader *bufio.Reader) (string, []byte, error) {
	//获取头部得到信息长度
	headLenByte, err := reader.Peek(4)
	if err != nil {
		return "", nil, err
	}
	//拿入缓冲区
	buffer := bytes.NewBuffer(headLenByte)
	//读取消息长度
	var dataLen int32
	err = binary.Read(buffer, binary.BigEndian, &dataLen)
	if err != nil {
		return "", nil, err
	}
	//判断长度
	if int32(reader.Buffered()) < dataLen+6 {
		fmt.Println("reader size < data size")
		return "", nil, err
	}
	dataByte := make([]byte, int(dataLen+6))
	//读取协议号
	_, err = reader.Read(dataByte[4:6])
	//读取data
	_, err = reader.Read(dataByte[6:])
	if err != nil {
		return "", nil, err
	}

	return string(dataByte[4:6]), dataByte[6:], nil
}
