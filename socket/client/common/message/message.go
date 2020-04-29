package message

import "fmt"

const(
	UserLogin = "用户登录成功"
	UserFail = "用户登录失败"
	UserUp = "用户上线"
	UserExit = "用户退出"
	NotUser = "用户不存在"
)

const (
	UserDown = iota  // 用户不在线
	UserIn  // 用户在线
)

type Message struct {
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
