package main

import "fmt"

type student struct{
	id int
	name string
	class string
}

func newStu(id int, name, class string) *student{
	return &student{
		id:id,
		name: name,
		class: class,
	}
}

type studentManage struct{
	stu []*student
}

func newstuma() *studentManage{
	return &studentManage{
		stu: make([]*student, 0, 50),
	}
}
// 增加
func (s *student) addmessage() {
	var (
		id int
		name string
		class string
	)	
	fmt.Scanf("id:%d", &id)
	fmt.Scanf("id:%d", &name)
	fmt.Scanf("id:%d", &class)
	newStu(id, name, class)
}
func (s *studentManage) addstu(st *student){
	st.addmessage()
	s.stu = append(s.stu, st)
}
// 删除
func (s *studentManage) substu(subid int){
	for i, v := range s.stu{
		if v.id == subid{
			s.stu = append(s.stu[:i], s.stu[:i+1]...)
			return
		}
	}
}
// 修改
func (s *studentManage) modify(st *student){
	for i, v := range s.stu{
		if v.id == st.id{
			s.stu[i] = st
			return
		}
	}
}
// 查找
func (s *studentManage) findstu(findid int) (a bool){
	for _, v := range s.stu{
		if v.id == findid{
			fmt.Printf("id:%d  name:%s   class:%s\n", v.id, v.name, v.class)
			a = true
		}else{
			fmt.Println("查无id")
			a = false
		}
	}
	return a
}
// 显示
func (s *studentManage) show(){
	for _, v := range s.stu{
		fmt.Printf("id:%d  name:%s   class:%s\n", v.id, v.name, v.class)
	}
}
