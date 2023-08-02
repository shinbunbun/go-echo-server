package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type StateMachine struct{}

var internalState int = 0

func (s *StateMachine) Add(input int, output *int) error {
	internalState += input
	*output = internalState
	return nil
}

func (s *StateMachine) Sub(input int, output *int) error {
	internalState -= input
	*output = internalState
	return nil
}

func (s *StateMachine) Mul(input int, output *int) error {
	internalState *= input
	*output = internalState
	return nil
}

func (s *StateMachine) Div(input int, output *int) error {
	if input == 0 {
		return errors.New("Divide by zero")
	}
	internalState /= input
	*output = internalState
	return nil
}

func main() {
	err := rpc.Register(new(StateMachine))
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
