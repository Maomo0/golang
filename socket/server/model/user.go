package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type User struct {
	Id int `json:"id"`
	Pwd string `json:"pwd"`
}

type UserInf interface {
	Encode() (data []byte, err error)
	Decode(data []byte) (err error)
	UserIn(r redis.Conn) (bool, []byte)
	GetUser(r redis.Conn) *User
}

func (u User) Encode() (data []byte, err error){
	data, err = json.Marshal(u)
	if err !=nil{
		fmt.Println("Encode user msg error", err)
		return
	}
	return data, err
}

func (u *User) Decode(data []byte) (err error){
	err = json.Unmarshal(data, &u)
	if err != nil{
		return err
	}
	return err
}

// 判断用户是否存在
func UserIn(r redis.Conn, id int) (bool, []byte){
	res, err := redis.String(r.Do("hget", "users", id))
	if err != nil{
		if err == redis.ErrNil{
			//fmt.Println(message.NotUser)
			return false, nil
		}
		return false, nil
	}
	return true, []byte(res)
}

func GetUser(r redis.Conn, id int) *User{
	u := &User{}
	b, res := UserIn(r, id)
	if !b{
		return nil
	}
	err := u.Decode(res)
	if err != nil{
		return nil
	}
	return u
}