package main

import (
	"errors"
	"fmt"
)

type person struct {
	name string
	age  uint
	address
}
type address struct {
	province string
	city     string
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Writer(p []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer
}

func main() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误：%w", e)
	fmt.Println(w)
	fmt.Println(errors.Is(w, e))
}
