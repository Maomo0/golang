package main

import "fmt"

type person struct{
	age int
	name string
	action string
}

type animal struct{
	name string
	action string
}

type inf interface {
	call() 
	move()
}

func (p *person) call()  {
	fmt.Printf("%d岁的%v在%v\n", p.age, p.name, p.action)
}

func (p *person) move() {
	fmt.Printf("%v在%v\n", p.name, p.action)
}

func (p *animal) call()  {
	fmt.Printf("一只%v在%v\n", p.name, p.action)
}

func (p *animal) move() {
	fmt.Printf("%v在%v\n", p.name, p.action)
}

func main(){
	var p inf
	p = &person{
		name: "她",
		age: 20,
		action: "洗头",
	}
	p.call()
	p.move()
	p = &animal{
		name: "狗狗",
		action: "活蹦乱跳",
	}
	p.call()
	p.move()
}