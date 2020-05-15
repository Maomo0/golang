package clientprocess

import (
	"bufio"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
	"os"
	"socket/client/cliutil"
	"socket/client/common/message"
	"socket/server/model"
	"strings"
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
	lo := model.LoginMsg{Id: id, LoginStatus: message.AfterLogin}
	switch str {
	case "1":
		lo.LoginChoice = str
		data := lo.Encode()
		err := cliutil.BufWrite(data, conn)
		if err != nil{
			return
		}
		ShowOnlineUser(id)
	case "2":
		// 信息的同步在message.LoginOnlinePro,协程开启能不断的接受服务其发送的信息
		lo.LoginChoice = str
		fmt.Printf("Sent data(输入exit退出发送):")
		for {
			data := bufio.NewReader(os.Stdin)   // 使用终端读取, scanf遇到空格读取不到
			str, _ = data.ReadString('\n')
			str = strings.Trim(str,"\t\n")
			if str == "exit"{
				break
			}
			lo.SentData = str
			da := lo.Encode()
			//str = fmt.Sprintf("sent msg to some people")
			err  := cliutil.BufWrite(da, conn)
			if err != nil{
				return
			}
		}
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
	us := &message.UserLoginMsg{}
	lo := &model.LoginMsg{}
	switch str {
	case "1":
		lo.LoginStatus = message.Login  // 发送给服务器选择的相关信息
		lo.LoginChoice = str
		data := lo.Encode()
		err := cliutil.BufWrite(data, conn)
		if err != nil{
			return err
		}
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
		lo.LoginStatus = message.Register
		lo.LoginChoice = str
		data := lo.Encode()
		err := cliutil.BufWrite(data, conn)
		if err != nil{
			return err
		}
		RegisterPro(conn)
		data, err = cliutil.ReadBuff(conn)
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
