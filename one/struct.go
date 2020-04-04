package main

import (
	"fmt"
	"os"
)

// type a struct {
// 	name string  //  `json:"title"`
// 	b  *b
// }
// type b struct {
// 	age int
// }
func showmanue(){
	fmt.Print()
	fmt.Println("1、增加信息")
	fmt.Println("2、删除信息")
	fmt.Println("3、修改信息")
	fmt.Println("4、查找信息")
	fmt.Println("0、显示信息")
}
func o(){
	sm := newstuma()
	stu1 := newStu(1, "1", "1")
	stu2 := newStu(2, "2", "2")
	stu3 := newStu(3, "3", "3")
	sm.stu = append(sm.stu, stu1, stu2, stu3)
	for{
		showmanue()
		var num int
		fmt.Scanf("输入数字:%d", &num)
		switch num{
			case 1: 
				var st student
				sm.addstu(&st)
			case 2: 
				var id int
				fmt.Scanf("输入删除id:%d", &id)
				sm.substu(id)
			case 3:
				var findid int 
				fmt.Scanf("输入需要修改id:%d", &findid)
				a := sm.findstu(findid)
				if a{
					var(
						id int
						name string
						class string
					)
					fmt.Scanf("新id", &id)
					fmt.Scanf("新name", &name)
					fmt.Scanf("新class", &class)
					stu := newStu(id, "name", "class")
					sm.modify(stu)
				}else {
					fmt.Println("输入id错误")
				}
				
			case 4:
				var findid int 
				fmt.Scanf("输入查找id:%d", findid)
				sm.findstu(findid)
			case 0: os.Exit(0)
		}
	}
}