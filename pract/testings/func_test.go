package main

import (
	_ "fmt"
	"log"
	"testing"
)
func TestLoopfunc(t *testing.T)  {
	res := loopfunc(5, 3)
	if res != 3{
		t.Fatalf("实际的值%d, 函数返回值%d", 3, res)
	}else{
		t.Log("正确")
		log.Println("log 正确")
	}
	
}

func TestForfunc(t *testing.T)  {
	res := forfunc(5, 3)
	if res != 3{
		t.Fatalf("实际的值%d, 函数返回值%d", 3, res)
	}else{
		t.Log("正确")
	}
}