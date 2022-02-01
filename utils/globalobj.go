package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

type Globalobj struct {
	TCPServer ziface.IServer
	Host string
	TCPPort uint32
	Name string

	Version string
	MaxConn int
	MaxSizePageSize uint32
}


var GlobalObject *Globalobj


func (g *Globalobj)Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &Globalobj{
		Name: "ZinServerApp",
		Version:"v0.4",
		TCPPort: 8999,
		Host: "0.0.0.0",
		MaxConn: 1000,
		MaxSizePageSize: 4096,
	}

	//GlobalObject.Reload()
}

