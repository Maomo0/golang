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
func ShowLoginMenu(r redis.Conn, conn net.Conn, id int){
	str,_ := cliutil.ReadMsg()  // 读取终端输入
	switch str {
	case "1":
		ShowOnlineUser(id)
		//if len(message.UsIn) == 1{
		//	fmt.Println(message.NotUserIn)
		//}
		//for _, v := range message.UsIn{
		//	if v.Id == id{
		//		continue
		//	}
		//	if v.UserStatus == message.UserIn{
		//		fmt.Printf("Id:%d  %s\n", v.Id, message.UserOnline)
		//	}
		//}
	case "2":
		str = fmt.Sprintf("sent msg to some people")
		err  := cliutil.BufWrite([]byte(str), conn)
		if err != nil{
			return
		}
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
func ChoiceMenu(r redis.Conn, conn net.Conn) error{
	str,_ := cliutil.ReadMsg()  // 读取终端输入
	err := cliutil.BufWrite([]byte(str), conn)
	us := &message.UserLoginMsg{}
	if err != nil{
		return err
	}
	switch str {
	case "1":
		id, pwd := message.LoginMsg()
		mes, ul, err := Login(conn, r, id, pwd)
		if err != nil{
			fmt.Println(err)
			return err
		}
		if mes.Type != message.UserIn {
			return err
		}
		us.Id = ul.Id
		us.UserStatus = ul.UserStatus
		us.Conn = conn
		fmt.Println("Id:", ul.Id, message.UserLogin)
		go func() {  // 协成开启不断接受服务端发送数据
			for{
				message.LoginOnlinePro(conn)
			}
		}()
		message.UpdateOnline(us, id, conn)
		for {
			LoginMenu()
			ShowLoginMenu(r, conn, us.Id)
		}
	case "2":
		RegisterPro(conn)
		data, err := cliutil.ReadBuff(conn)
		if err != nil{
			break
		}
		if string(data) == message.UserIns || string(data) == message.UserRegister{
			fmt.Println(string(data))
			break
		}
	case "3":
		os.Exit(0)
	default:
		fmt.Println("无该选项")
	}
	return nil
}
