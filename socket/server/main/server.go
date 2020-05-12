package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net"
	"socket/server/serutil"
	serpro "socket/server/serverprocess"
)

var(
	P serutil.RedisConnect
	list net.Listener
	err error
	conn net.Conn
)

func init(){
	P = serutil.InitRedis(":6379")
	list, err = net.Listen("tcp", ":8999")
	if err != nil{
		log.Fatal("Start the server fail\n", err)
	}

}
func closes(conn net.Conn, list net.Listener, rc redis.Conn){
	func() {
		err = list.Close()
		if err != nil {
			log.Fatal("close the list connection fail\n", err)  // log.fatal 执行后立即结束
		}
		err = conn.Close()
		if err != nil{
			log.Fatalln("Can't close conn")
		}
		err = rc.Close()
		if err != nil{
			log.Println("redis can't stop")
		}
	}()
}
func main(){
	rc := P.Pool.Get()
	var msg []byte
	defer closes(conn, list, rc)
	fmt.Println("server start successes wait for a client connect")
	for  {
		conn, err = list.Accept()
		//var b = true
		if err != nil{
			log.Fatal("connect fail\n", err)
		}
		log.Println("connect addr for", conn.RemoteAddr())
		go func() {
			for{
				msg, err, conn = serutil.ReadBuff(conn)  // 读取客户端终端输入
				serpro.Main(string(msg), conn, err, rc)
				fmt.Println(string(msg))
				//for _ , v := range serpro.Usg.GetOnlineUser(){
				//	fmt.Println(v.Conn.RemoteAddr())
				//}
				if err != nil{
					serpro.UserExit(conn)
					//for _ , v := range serpro.Usg.GetOnlineUser(){
					//	fmt.Println("server 65", v.Conn.RemoteAddr())
					//}
					break
				}
					//serpro.Main(string(msg), conn)
					//serpro.SentOnline(conn)
			}
		}()
	}
}
