package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	print("Please input message: ")
	var input string
	fmt.Scanln(&input)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("EchoService.Echo", input, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}
