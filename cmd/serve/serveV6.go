package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouterV6 struct {
	znet.BaseRouter
}


func (pr *PingRouterV6) Handle(request ziface.IRequest) {
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

type HelloZinRouter struct {
	znet.BaseRouter
}

func (this *HelloZinRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinRouter Handle")
	fmt.Println("recv from client: msgId=", request.GetMsgID(), ",data =", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("hello zinx router v0.6"))
	if err != nil {
		fmt.Println(err)
	}
}


func main() {
	s := znet.NewServe()

	s.AddRouter(0, &PingRouterV6{})
	s.AddRouter(1, &HelloZinRouter{})
	s.Serve()
}
