package main

import (
	"fmt"
	"strings"
)
const (
	_  = iota
	n1 = 1 << (10 * iota)
	n2 = 1 << (10 * iota)
	n3 = 1 << (10 * iota)
	n4 = 1 << (10 * iota)
)

func x() {
	fmt.Println(n1, n2, n3, n4)
	// s := "heloo 哈哈"
	// for i := 0; i < len(s); i++{
	// 	fmt.Println(string(s[i]))
	// }
	// for _, j := range s{
	// 	fmt.Println(string(j))
	// }
	a := "string"
	fmt.Println(strings.Split(a, "")[1])
	fmt.Println(a)
	s := [] string{"hello", "哈哈"}
	fmt.Println(s)

	var b int
	fmt.Scanf("%d", &b)
	fmt.Println(b)
}
