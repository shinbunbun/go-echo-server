package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type EchoService struct{}

func (s *EchoService) Echo(input string, output *string) error {
	*output = input
	return nil
}

func main() {
	print("Please input message: ")
	var input string
	fmt.Scanln(&input)
	err := rpc.Register(new(EchoService))
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
