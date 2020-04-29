package message

import "net"

var onlineUser map[int]*UserMsg

func init(){
	onlineUser = make(map[int]*UserMsg, 1024)
}
type Users struct {
	Id int `json:"id"`
	Pwd string `json:"pwd"`
	Status string `json:"status"`
}
type UserMsg struct {
	Id int
	Conn net.Conn
}