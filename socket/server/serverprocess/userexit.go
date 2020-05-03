package serverprocess

import (
	"fmt"
	"net"
)

func UserExit(conn net.Conn){
	id := GetOnlineId(conn)
	if id != 0{
		Usg.DelOnlineUer(id)
	}else {
		fmt.Println("用户无法删除")
	}
}
