package main

import (
	"log"
	"time"
)

type serveOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOption() *serveOption {
	return &serveOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serveOption) {}

func main() {
	opt := newOption()
	opt.port = 8085
	server(opt)
}
