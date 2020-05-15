package serverprocess

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"socket/client/common/message"
	"socket/server/serutil"
)


func Main(mes string, conn net.Conn, err error, r redis.Conn)  {
	switch mes {
	case "1":
		ManOne(conn, err)
	case "2":
		ChoiceTwo(r, conn)
	case "3":
		err = conn.Close()
		if err != nil{
			return
		}
	}
}
func ManOne(conn net.Conn, err error)  {
	ul, id := LoginPro(conn)
	if err != nil{
		fmt.Println(err)
		if id != 0{
			fmt.Printf("%d", id)
			Usg.DelOnlineUer(id)
		}
		return
	}
	switch ul.UserStatus {
	case message.UserIn:
		for _, v := range Usg.GetOnlineUser() {
			SentOnline(v.Conn)
		}
	}
}
func ChoiceTwo(r redis.Conn, conn net.Conn){
	b := RegisterPro(conn, r)
	if b{
		serutil.BufWrite([]byte(message.UserRegister), conn)
	}else{
		serutil.BufWrite([]byte(message.UserIns),conn)
	}
}