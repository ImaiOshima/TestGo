package main

import (
	"fmt"
	"unsafe"
)

func toString(byte []byte) string {
	return *(*string)(unsafe.Pointer(&byte))
}

func main() {
	bs := []byte("hello world")
	s := toString(bs)
	fmt.Println(s)
	fmt.Printf("bs - %x\n", &bs)
	fmt.Printf("s - %x\n", &s)
}
