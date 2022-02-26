package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 1
	fmt.Println("a:", unsafe.Sizeof(a))
	b := int32(1)
	fmt.Println("b:", unsafe.Sizeof(b))
	len := int32(len("123456"))
	fmt.Println("len:", len)
	fmt.Println("b:", unsafe.Sizeof(len))
}
