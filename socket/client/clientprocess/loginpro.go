package clientprocess

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"socket/client/cliutil"
	"socket/client/common/message"
	"socket/server/model"
	serpro "socket/server/serverprocess"
)

func Login(conn net.Conn, r redis.Conn, id int, pwd string) (*message.Message, *serpro.UserLoginMsg, error){
	u := &model.User{}
	mes := &message.Message{}
	ul := &serpro.UserLoginMsg{}
	//var useronline = &message.UserOnlineList{}
	u.Id = id
	u.Pwd = pwd
	l := model.GetUser(r, u.Id)  // 获取用户
	if l == nil{   // 用户不存在
		err := errors.New(message.NotUser)
		return mes, ul, err
	}
	data, err := u.Encode()
	if err != nil{
		return mes, ul, err
	}
	if u.Pwd != l.Pwd{
		err = errors.New(message.UserFail)
		return mes, ul, err
	}
	// 登录成功后
	mes.Data = string(data)
	mes.Type = message.UserIn
	data,_ = mes.Encode()
	err = cliutil.BufWrite(data, conn)   // 向服务器发送客户端输入的登录信息
	if err != nil{
		return mes, ul, err
	}
	ul.Conn = conn
	ul.Id = u.Id
	ul.UserStatus = message.UserIn
	return mes, ul, nil
}

func ShowOnlineUser(id int){
	// 用户登录后选择1,查看当前在线人数
	if len(message.UsIn) == 1{
		fmt.Println(message.NotUserIn)
	}
	for _, v := range message.UsIn{
		if v.Id == id{
			continue
		}
		if v.UserStatus == message.UserIn{
			fmt.Printf("Id:%d  %s\n", v.Id, message.UserOnline)
		}
	}
}

func SentMsgToOnlineUser(me * message.Message) *message.Message{
	return me
}

func AddFriend(){
	//用户登录后选择3,添加好友
}

func SentMsgToFriend()  {
	//用户登录后选择4,向好友发送信息
	//需要维护一个好友列表,选择时显示
}
