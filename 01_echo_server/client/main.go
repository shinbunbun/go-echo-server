package main

import (
	"fmt"
	"net"
)

func main() {
	print("Please input message: ")
	var input string
	fmt.Scanln(&input)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte(input))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))
}
