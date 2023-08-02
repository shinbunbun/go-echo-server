package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	print("Please input command: ")
	var order string
	var value int
	fmt.Scan(&order, &value)
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int
	switch order {
	case "add":
		err = client.Call("StateMachine.Add", value, &reply)
	case "sub":
		err = client.Call("StateMachine.Sub", value, &reply)
	case "mul":
		err = client.Call("StateMachine.Mul", value, &reply)
	case "div":
		err = client.Call("StateMachine.Div", value, &reply)
	}
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(reply)
}
