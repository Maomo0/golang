package cliutil

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Connection struct {
	Conn net.Conn
	UserId int
}
func CliClose(conn net.Conn)  {
	err := conn.Close()
	if err != nil{
		fmt.Errorf("%s", err)
		return
	}
}
func InitConn () (cn Connection, err error){
	cn.Conn, err = net.Dial("tcp", ":8999")  // 注意:
	if err != nil{
		log.Fatalln("Connecting the server fail", err)
		return
	}
	return cn, err
}
// 读取终端输入信息
func ReadMsg() (str string, err error){
	fmt.Println("please choice one digital:")
	data := bufio.NewReader(os.Stdin)
	str, err = data.ReadString('\n')
	if err != nil{
		fmt.Println("Can't read data", err)
		return
	}
	str = strings.Trim(str,"\t\n")
	return str, err
}
func ReadBuff(conn net.Conn)(msg []byte, err error){
	//defer CliClose(conn)  // 确保了只读一次
	b := bufio.NewReader(conn)
	lengthByte, _ := b.Peek(4)   // 读取前4个字节,根据长度的数据类型而定如int32为4
	lenBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lenBuff, binary.LittleEndian, &length)
	if err != nil{
		log.Fatalln("clii read err cliutil/client 61", err)
		return nil, err
	}
	msg = make([]byte, int(4+length))
	_, err = b.Read(msg)
	if err != nil{
		log.Fatalln("cli read err main/client 66", err)
		return nil, err
	}
	return msg[4:], err
}

func BufWrite(msg []byte, conn net.Conn) (err error){
	//defer CliClose(conn)
	buff := new(bytes.Buffer)   // 创建缓存对象
	length := int32(len(msg))	// int32的长度
	err = binary.Write(buff, binary.LittleEndian, length)
	if err != nil{
		log.Fatalln("cli sent err main/client 63", err)
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, msg)   // big大端模式地址从小到大, list小端模式地址从大到小0x78,0x56, 一般用小端
	if err != nil{
		log.Fatalln("binary write err main/client 68", err)
		return err
	}
	_, err = conn.Write(buff.Bytes())
	return err
}
