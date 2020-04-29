package clientprocess

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"os"
	"socket/client/cliutil"
	"socket/client/common/message"
)

func MainMenu() {
	fmt.Println("\t=================\t")
	fmt.Println("\t1:用户登录\t")
	fmt.Println("\t2:用户注册\t")
	fmt.Println("\t3:退出\t")
	fmt.Println("\t=================\t")
}
func LoginMenu()  {
	fmt.Println("\t-------------------\t")
	fmt.Println("\t1:查询在线人数\t")
	fmt.Println("\t2:群发消息\t")
	fmt.Println("\t3:添加好友\t")
	fmt.Println("\t4:给好友发送信息\t")
	fmt.Println("\t5:用户退出\t")
	fmt.Println("\t-------------------\t")
}
func ShowLoginMenu(r redis.Conn, conn net.Conn){
	str,_ := cliutil.ReadMsg()  // 读取终端输入
	switch str {
	case "1":
		fmt.Println("sum person")
	case "2":
		fmt.Println("sent msg to some people")
	case "3":
		fmt.Println("add one friend")
	case "4":
		fmt.Println("sent msg to your friend")
	case "5":
		os.Exit(0)
	default:
		fmt.Println("无该选项")
	}
}
func ChoiceMenu(r redis.Conn, conn net.Conn){
	str,_ := cliutil.ReadMsg()  // 读取终端输入
	err := cliutil.BufWrite([]byte(str), conn)
	var msg []byte
	if err != nil{
		return
	}
	switch str {
	case "1":
		id, pwd := message.LoginMsg()
		mes, _ := Login(conn, r, id, pwd)
		if mes.Type != message.UserIn {
			break
		}
		msg, err = cliutil.ReadBuff(conn)
		if err != nil{
			return
		}else {
			fmt.Println(string(msg))
		}
		for {
			LoginMenu()
			ShowLoginMenu(r, conn)
		}
	case "2":
		fmt.Println("register")
	case "3":
		os.Exit(0)
	default:
		fmt.Println("无该选项")
	}
}