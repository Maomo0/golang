package serverprocess

import (
	"fmt"
	"net"
	"socket/client/common/message"
	"socket/server/serutil"
)


func Main(mes string, conn net.Conn)  {
	switch mes {
	case "1":
		ManOne(conn)
	case "2":
		fmt.Println("2")
	case "3":
		err := conn.Close()
		if err != nil{
			return
		}
	}
}
func ManOne(conn net.Conn)  {
	LoginPro(conn)
	err := serutil.BufWrite([]byte(message.UserUp), conn)
	if err != nil{
		return
	}
}