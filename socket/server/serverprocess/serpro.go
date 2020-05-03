package serverprocess

import (
	"fmt"
	"net"
	"socket/client/common/message"
)


func Main(mes string, conn net.Conn, err error)  {
	switch mes {
	case "1":
		ManOne(conn, err)
	case "2":
		fmt.Println("2")
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
	//err := serutil.BufWrite([]byte(message.UserUp), conn)
	//if err != nil{
	//	return
	//}
}