package serverprocess

import (
	"encoding/json"
	"fmt"
	"net"
	"socket/client/common/message"
	"socket/server/model"
	"socket/server/serutil"
)

type UserLoginMsg struct {
	 Id int `json:"id"`
	 UserStatus int `json:"user_status"`
	 Conn net.Conn   `json:"conn"`
}
// 读取登录信息
func LoginPro(conn net.Conn) (*UserLoginMsg, int){
	u := &model.User{}
	mes := &message.Message{}
	ul := &UserLoginMsg{}
	data, err ,_:= serutil.ReadBuff(conn)
	if err != nil{
		return ul, 0
	}
	mes, _ = mes.Decode(data)
	err = u.Decode([]byte(mes.Data))
	if err != nil{
		return ul, 0
	}
	ul.Conn = conn
	ul.Id = u.Id
	ul.UserStatus = message.UserIn
	Usg.AddOnlineUser(ul)
	fmt.Println("read data:", u)
	return ul, ul.Id
}

func SentOnline(conn net.Conn){
	ut := Usg.GetOnlineUser()
	data, _ := json.Marshal(ut)
	err :=serutil.BufWrite(data, conn)
	if err != nil{
		return
	}
}
