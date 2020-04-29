package function

import (
	"encoding/json"
	"fmt"
	"net"
	"socket/server/model"
	"websocket/client/util"
)

func Login(conn net.Conn) (data []byte){
	var u model.User
	fmt.Scanf("Id:%d", &u.Id)
	fmt.Scanf("Pwd:%s", &u.Pwd)
	data, _ = json.Marshal(u)
	err := util.BufWrite(data, conn)
	if err != nil{
		return
	}
	return  data
}