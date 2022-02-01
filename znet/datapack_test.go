package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("server listen error :", err)
	}

	go func() {
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accpet error")
				return
			}

			go func(conn net.Conn) {
				dp := NewDataPack()

				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)

					if err != nil {
						fmt.Println("read head error")
						return
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error")
						return
					}

					if msgHead.GetMsgLen() > 0 {
						//进行二次读取
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err:", err)
							return
						}

						fmt.Println("--->resv msgID:", msg.Id,", datalen=", msg.DataLen, "data = ", string(msg.Data))

					}
				}

			}(conn)
		}
	}()


	/*
	模拟客户端请求
	 */

	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}
	dp := NewDataPack()
	//模拟粘包过程

	msg1 := &Message{
		Id: 1,
		DataLen: 4,
		Data: []byte{'z','i','n','x'},
	}

	sendData1, err := dp.Pack(msg1)

	if err != nil {
		fmt.Println("client pack error", err)
	}

	msg2 := &Message{
		Id: 1,
		DataLen: 7,
		Data: []byte{'o','h','o','h','o','o','0'},
	}

	sendData2, err := dp.Pack(msg2)

	if err != nil {
		fmt.Println("client pack error", err)
	}
	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	select{}

}
