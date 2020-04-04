package main

import (
	"encoding/json"
	"fmt"
	iou "io/ioutil"
)

type stu struct {
	Name string
	Skill string
	Age int
}

type student interface {
	Store() bool
	Restore() (bool, []byte)
}

func (s *stu) Store() bool {
	data, err := json.Marshal(s)
	if err != nil{
		fmt.Println("序列化错误", err)
		return false
	}
	errt := iou.WriteFile("./json.json", data, 0777)
	if errt != nil{
		fmt.Println("写入错误", errt)
		return false
	}
	return true
}

func (s *stu) Restore() (bool, []byte)  {
	data, err := iou.ReadFile("./json.json")
	if err != nil{
		fmt.Println("读入错误", err)
		return false, nil
	}
	errt := json.Unmarshal(data, s)
	if errt != nil{
		fmt.Println("反序列化错误", errt)
		return false, nil
	}
	return true, data
}