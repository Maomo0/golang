package serverprocess

import (
	"log"
	"net"
	"socket/client/common/message"
)

var Usg *UserMsg
type UserMsg struct {
	onlineUser map[int]*UserLoginMsg
}
func init()  {
	Usg = &UserMsg{
		onlineUser: make(map[int]*UserLoginMsg, 1024),
	}
}
// 添加在线用户
func (u *UserMsg) AddOnlineUser(ul *UserLoginMsg){
	u.onlineUser[ul.Id] = ul
}
// 删除
func (u *UserMsg) DelOnlineUer(id int){
	delete(u.onlineUser, id)
}
// 返回当前在线用户
func (u *UserMsg) GetOnlineUser() map[int]*UserLoginMsg{
	return u.onlineUser
}
func GetOnlineId(conn net.Conn) (id int){
	for _, v := range Usg.onlineUser{
		if v.Conn.RemoteAddr() == conn.RemoteAddr(){
			log.Printf("Id:%d  %s\n", v.Id,message.UserExit)
			return v.Id
		}
	}
	return 0
}
