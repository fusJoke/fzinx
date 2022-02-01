package main

import "zinx/znet"


func main() {
	s := znet.NewServe("[zinx V0.1]")
	s.Serve()
}
