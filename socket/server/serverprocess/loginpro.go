package serverprocess

import (
	"fmt"
	"net"
	"socket/server/model"
	"socket/server/serutil"
)

type UserLoginMsg struct {
	 Conn net.Conn
	 Id int
}
// 读取登录信息
func LoginPro(conn net.Conn){
	u := &model.User{}
	data, _ := serutil.ReadBuff(conn)
	err := u.Decode(data)
	if err != nil{
		return
	}
	ul := &UserLoginMsg{
		Conn: conn,
		Id:   u.Id,
	}
	Usg.AddOnlineUser(ul)
	fmt.Println("read data:", u)
}
