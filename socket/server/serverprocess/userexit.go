package serverprocess

import (
	"net"
	"socket/client/common/message"
)

func UserExit(conn net.Conn){
	id := GetOnlineId(conn)
	SetExitMsg(id)
	if id != 0{
		for _, v := range Usg.onlineUser{
			if v.Id != id{
				SentOnline(v.Conn)  // 向在线用户发送信息,包括退出用户的信息
			}
			//fmt.Println(v.Id, v.Conn.RemoteAddr())
		}
		Usg.DelOnlineUer(id)  // 服务器删除在线用户
	}else {
		return
	}
}
func SetExitMsg(id int){
	u := Usg.GetOnlineUser()
	for _, v := range u{
		if v.Id == id{
			v.UserStatus = message.UserDown
			break
		}
	}
}