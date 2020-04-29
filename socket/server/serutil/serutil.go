package serutil

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net"
)
type RedisConnect struct {
	Pool *redis.Pool
}
func InitRedis(port string) (p RedisConnect){
	p.Pool = &redis.Pool{
		MaxActive: 3,
		MaxIdle: 4,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", port)
			if err != nil{
				log.Fatal("Can't connect redis server", err)
			}
			return conn, err
		},
	}
	return p
}
func ReadBuff(conn net.Conn)(msg []byte, err error) {
	b := bufio.NewReader(conn)
	lengthByte, _ := b.Peek(4) // 读取前4个字节,根据长度的数据类型而定如int32为4
	lenBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err = binary.Read(lenBuff, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("server binary read err serutil/serutil 44", err, conn.RemoteAddr())
		return nil, err
	}
	msg = make([]byte, int(4+length))
	_, err = b.Read(msg)
	if err != nil {
		fmt.Println("server read err serutil/serutil 50", err, conn.RemoteAddr())
		return nil, err
	}
	return msg[4:], err
}

func BufWrite(msg []byte, conn net.Conn) (err error){
	buff := new(bytes.Buffer)   // 创建缓存对象
	length := int32(len(msg))	// int32的长度
	err = binary.Write(buff, binary.LittleEndian, length)
	if err != nil{
		fmt.Println(".....")
		return err
	}
	err = binary.Write(buff, binary.LittleEndian, msg)   // big大端模式地址从小到大, list小端模式地址从大到小0x78,0x56, 一般用小端
	if err != nil{
		fmt.Println("数据")
		return err
	}
	_, err = conn.Write(buff.Bytes())
	return err
}

