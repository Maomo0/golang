package main

import "fmt"

func sayName(name string){
	fmt.Println("hello", name)
}

func deferTest() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func sum(x, y int) int{
	xy := x+y
	return xy
}

func sub(x, y int) (xy int){
	xy = x-y
	return
}

func all(x, y int, op func(int, int) int) int{
	return op(x, y)
}
// 可变参数
func lot(x int, a... int){
	fmt.Println(x)
	for _, v := range a{
		fmt.Println(v)
	}
}

// 闭包
func closefunc(name string) func(){
	return func(){
		fmt.Println("welocme", name)
	}
}

func numall(x, y int) (func() int, func() int){
	subxy := func() int{
		xy := x-y
		return xy
	}
	sumxy := func() int{
		xy := x+y
		return xy 
	}
	return subxy, sumxy
}

// 错误
func erroruse() {
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("错误")
		}
	}()
	panic("抛出")
}

// 自定义类型、结构体
type  myint int

type user struct {
	id int 
	name string
}

// 方法
type person struct{
	name string
	age int
}

func newPerson(name string, age int) *person{
	// 构造函数
	return &person{
		name:name,
		age:age,
	}
}

func (p person) sayname(){
	// 方法
	fmt.Println("your age", p.name)
}

func (p *person) setage(age int){
	// 改变年龄
	p.age = age
}

// 接口
// type dog struct {
// 	name string
// }
// type doger interferce{
// 	say()
// }
// func (d dog) say(){
// 	fmt.Println("www")
// }
func tee(){
	// fmt.Println("start")
	// deferTest()
	// fmt.Println("end")

	// a :=  closefunc("go")
	// a()

	// x := all(100, 200, sum)
	// fmt.Println(x)
	// sub, sum := numall(100, 200)
	// fmt.Println(sub(), sum())
	//lot(1, 2, 3, 4)
	
	//erroruse()
	// var b myint
	// b = 10
	// fmt.Println(b)
	// a := user {
	// 	id:1 ,
	// 	name:"hh",
	// }
	// fmt.Println(a)

	// a := newPerson("miao", 29)
	// fmt.Println(a.name)
	// a.setage(35)
	// fmt.Println(a.age)
}