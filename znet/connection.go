package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	ExitChan chan bool
	MsgHandler ziface.IMsgHandle
	msgChan chan []byte
}

func NewConnection(c *net.TCPConn,id uint32, MsgHandler ziface.IMsgHandle ) *Connection {
	return &Connection{
		Conn:c,
		ConnID: id,
		isClosed: false,
		ExitChan:make(chan bool, 1),
		MsgHandler: MsgHandler,
		msgChan:make(chan []byte),
	}
}

func(c *Connection) StartReader() {
	fmt.Println("Reader goruotine is running ")
	defer fmt.Println("connId =", c.ConnID, "reader is exit, remoteAddr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//buf := make([]byte, utils.GlobalObject.MaxSizePageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error", err)
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			io.ReadFull(c.GetTCPConnection(), data)
		}
		msg.SetData(data)

		req := &Request{
			conn:c,
			msg:msg,
		}
		go c.MsgHandler.DoMsgHandler(req)
	}
}

func(c *Connection) StartWrite() {
	fmt.Println("[write Goroutine is running ] ")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit]")

	for {
		select {
			case data := <- c.msgChan:
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Data error :", err ," conn writer exit")
					return
				}
			case <- c.ExitChan:
				return
		}
	}
}



func(c *Connection) Start() {
	fmt.Println("Conn start() ... ConnID = ", c.ConnID)
	go c.StartReader()
	go c.StartWrite()

	for {
		select {
			case <- c.ExitChan:
				return 
		}
	}
}
func(c *Connection) Stop() {
	fmt.Println("Conn stop() .... ConnID", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true

	c.Conn.Close()

	close(c.ExitChan)

}
func(c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func(c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func(c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func(c *Connection) AddRouter(msgID uint32, router ziface.IRouter) {
	c.MsgHandler.AddRouter(msgID, router)
}

func(c *Connection) Send([]byte) error {
	return errors.New("eccc")
}


func(c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return  errors.New("connection closed erro")
	}

	dp := NewDataPack()

	Msg, err := dp.Pack(NewMsgPackage(msgId, data))

	if err != nil {
		fmt.Println("pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}

	c.msgChan <- Msg

	return nil
}
