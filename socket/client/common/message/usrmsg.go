package message

import (
	"encoding/json"
	"fmt"
	"net"
	"socket/client/cliutil"
)

var UsIn map[int]*UserLoginMsg // 好友数
type UserLoginMsg struct {
	Id int `json:"id"`
	UserStatus int`json:"user_status"`
	Conn net.Conn   `json:"conn"`
}
func init(){
	UsIn = make(map[int]*UserLoginMsg, 1024)
}
func LoginOnlinePro(conn net.Conn){
	// 读取服务器发送数据
	msg ,err := cliutil.ReadBuff(conn)
	if err != nil{
		fmt.Println("can't read server data")
	}
	err  = json.Unmarshal(msg, &UsIn)
	for _, v := range UsIn{   // 检查是否有用户退出
		if v.UserStatus == UserDown{
			delete(UsIn, v.Id)
			break
		}
	}
	if err != nil{
		return
	}
}
func PrintOnline(id int){
	for _,v := range UsIn{
		if v.Id == id{
			continue
		}
		str := fmt.Sprintf("Id:%d %s\n", v.Id, UserUp)
		fmt.Printf(str)
	}
}
func UpdateOnline(um *UserLoginMsg, id int, conn net.Conn){
	// 更新在线信息
	ut , ok := UsIn[id]
	if !ok {
		ut = &UserLoginMsg{
			Id: id,
			UserStatus: UserIn,
			Conn: conn,
		}
	}
	ut.Conn = conn
	UsIn[um.Id] = ut
	//PrintOnline(ut.Id)
}