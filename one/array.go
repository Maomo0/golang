package main

import (
	"fmt"
	"sort"
)

func te(){
	// a := [...]string{1:"python", 2:"c++", 3:"goland"}
	// for _, value := range a{
	// 	fmt.Println(value)
	// }

	// for index, value := range b{
	// 	for i, j := range b{
	// 		if i == index{
	// 			continue
	// 		}
	// 		if value + j == 8{
	// 			fmt.Printf("(%d, %d)", i, index)
	// 		}
	// 	}
	// }
	a := make([]string, 5)
	b := []int{1, 3, 5, 7, 8}
	c := []string{"北", "上", "广", "深"} 
	
	d := fmt.Sprintf("%d %d", b[1], b[2])
	a = append(a, c[:]...)
	fmt.Println(a[0:4])
	sort.Ints(b[:])  // 转换为切片
	c = append(c[0:2], c[3:]...)
	b = append(b[0:2], b[3:]...)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}