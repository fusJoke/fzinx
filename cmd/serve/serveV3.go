package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call Router Prehandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before bing"))

	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Router handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("handle bing"))

	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call Router after handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after bing"))

	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func main() {
	s := znet.NewServe()

	s.AddRouter(&PingRouter{})
	s.Serve()
}
