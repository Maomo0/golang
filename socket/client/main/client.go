package main

import (
	pro "socket/client/clientprocess"
	"socket/client/cliutil"
	"socket/server/serutil"
)

var c cliutil.Connection
var p serutil.RedisConnect
func init(){
	c, _ = cliutil.InitConn()
	p = serutil.InitRedis(":6379")
}

func main(){
	defer cliutil.CliClose(c.Conn)
	for  {
		pro.MainMenu()
		pro.ChoiceMenu(p.Pool.Get(), c.Conn)
	}
}
