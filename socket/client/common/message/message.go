package message

import (
	"encoding/json"
	"fmt"
)

const(
	UserLogin = "用户登录成功"
	UserFail = "密码错误"
	UserUp = "用户上线"
	UserExit = "用户退出"
	NotUser = "该用户不存在"
	UserOnline = "用户在线"
	UserIns = "用户已存在"
	UserRegister = "用户注册成功"
	NotUserIn = "无用户在线"
)
const (
	AfterLogin = "AfterLogin"
	Login = "Login"
	Register = "Register"
)
const (
	UserDown = iota  // 用户不在线
	UserIn  // 用户在线
	SentData
)

type Message struct {
	Id int `json:"id"`
	Data string `json:"data"`
	Type int `json:"type"`
}
// 记录登录的信息
func LoginMsg() (id int, pwd string){
	fmt.Printf("Id:")
	fmt.Scanf("%d\n", &id)
	fmt.Printf("Pwd:")
	fmt.Scanf("%s\n", &pwd)
	return id, pwd
}

func (m *Message) Encode()(data []byte, err error){
	data, err = json.Marshal(m)
	if err != nil{
		return
	}
	return data, nil
}
func (m *Message) Decode(data []byte) (me *Message, err error) {
	err = json.Unmarshal(data, &m)
	if err != nil{
		return
	}
	return m, nil
}