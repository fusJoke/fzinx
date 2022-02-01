package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")

	time.Sleep(1*time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil{
		fmt.Println("net dial failed :", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello znx v0.1"))

		if err != nil {
			fmt.Println("conn write failed :", err)
			return
		}

		buf := make([]byte, 512)

		cnt, err := conn.Read(buf)

		if err != nil {
			fmt.Println("read buf failed :", err)
			return
		}

		fmt.Printf("serve back:%s, cnt = %d\n", buf, cnt)

		time.Sleep(1*time.Second)
	}

}
