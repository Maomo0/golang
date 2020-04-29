package clientprocess

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"socket/client/common/message"
	"socket/server/model"
	serpro "socket/server/serverprocess"
	"websocket/client/util"
)

func Login(conn net.Conn, r redis.Conn, id int, pwd string) (*message.Message, *serpro.UserLoginMsg){
	u := &model.User{}
	mes := &message.Message{}
	ul := &serpro.UserLoginMsg{}
	u.Id = id
	u.Pwd = pwd
	l := model.GetUser(r, u.Id)  // 获取用户
	if l == nil{   // 用户不存在
		return nil, nil
	}
	data, err := u.Encode()
	if err != nil{
		return nil, nil
	}
	err = util.BufWrite(data, conn)   // 向服务器发送客户端输入的登录信息
	if err != nil{
		return nil, nil
	}
	if u.Pwd != l.Pwd{
		fmt.Println("密码错误")
		return nil, nil
	}
	// 登录成功后
	mes.Data = string(data)
	mes.Type = message.UserIn
	ul.Conn = conn
	ul.Id = u.Id
	//serpro.Usg.AddOnlineUser(ul)
	return mes, ul
}

// 用户登录后选择1
func LoginOne(conn net.Conn){

}
