package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port uint32
	MsgHandler ziface.IMsgHandle
}

func(s *Server) Start(){

	fmt.Printf("[zinx] server name: %s, listener at ip: %s, port : %d\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host,utils.GlobalObject.TCPPort)
	fmt.Printf("[zinx] version %s, MaxConn:%d, MaxPackageSize:%d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxSizePageSize)
	fmt.Printf("[Start] Server Listener at IP:%s, Port: %d, is starting\n",s.IP, s.Port)

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP,s.Port))

		if err != nil {
			fmt.Println("resolve addr error:", err)
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)

		if err != nil {
			fmt.Println("listen", s.IPVersion, "err:", addr)
		}

		var cid uint32
		cid = 0

		fmt.Println("start zinx serve success,", s.Name, "listening..." )

		for {
			conn, err := listenner.AcceptTCP()

			if err != nil {
				fmt.Println("Accept err", err)
			}

			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			go dealConn.Start()
		}
	}()

}


func(s *Server) Stop(){

}

func(s *Server) Serve(){
	s.Start()

	select {}
}


func NewServe() ziface.IServer {
	s := &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TCPPort,
		MsgHandler: NewMsgHandler(),
	}

	return s
}

func (s *Server) AddRouter (msgID uint32, router ziface.IRouter)  {
	s.MsgHandler.AddRouter(msgID, router)
}
