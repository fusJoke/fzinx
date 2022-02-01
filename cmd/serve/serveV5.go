package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouterV5 struct {
	znet.BaseRouter
}


func (pr *PingRouterV5) Handle(request ziface.IRequest) {
	fmt.Println("call Router handle")

	fmt.Println("recv from client: msgId = ",
		request.GetMsgID(),
		", data =", string(request.GetData()))

	err := request.GetConnection().SendMsg( 1, []byte("ping...ping...ping"))

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("call back before ping error")
	}
}


func main() {
	s := znet.NewServe()

	s.AddRouter(&PingRouterV5{})
	s.Serve()
}
