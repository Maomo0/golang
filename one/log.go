package main

import (
	// "database/sql"
	"fmt"
	"log"
	//"strconv"
	"time"
)

func main()  {
	c := []int{1, 2, 3}
	for i:=0 ; i<2; i++{
		log.Println("测试日志")
		time.Sleep(time.Second*3)
	}
	fmt.Println(1<<2)
	a := "你是个"
	b := "傻子"
	fmt.Println(append(c[:0], c[1:]...))
	fmt.Println(a+b)
}