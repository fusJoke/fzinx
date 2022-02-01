package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
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

		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("zinx0.5 client test Message")))
		if err != nil {
			fmt.Println("pack error:", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err !=nil {
			fmt.Println("write error", err)
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}
		msgHead, err := dp.Unpack(binaryHead)

		if err != nil {
			fmt.Println("client unpack msgHead error", err)
		}

		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error, ", err)
				return
			}
			fmt.Println("------> Recv Server Msg : ID =", msg.Id,
				"len=", msg.DataLen,
				", data =", string(msg.Data),
				)
		}
		time.Sleep(1*time.Second)
	}

}
