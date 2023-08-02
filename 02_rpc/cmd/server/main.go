package main

import (
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
	err := rpc.Register(new(EchoService))
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal(err)
	}
}
