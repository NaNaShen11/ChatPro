package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type host struct {
	Host string
}

var Conf = host{}
//初始化配置
func init() {
	loadFile,err :=os.Open("src/config/config.json")
	if err!=nil{
		fmt.Println("config load err....")
		return
	}
	defer loadFile.Close()
	decode := json.NewDecoder(loadFile)
	err = decode.Decode(&Conf)
	if err!=nil{
		fmt.Println("config decode err...")
		return
	}
}