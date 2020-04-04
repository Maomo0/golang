package main

import (
	"fmt"
	"runtime"
)

func main()  {
	a := runtime.NumCPU()
	fmt.Println(a)
}