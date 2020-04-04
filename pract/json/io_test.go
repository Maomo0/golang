package main

import "testing"

func TestStore(t *testing.T)  {
	var stuin student
	stuin = &stu{
		Name: "狗",
		Skill: "跳",
		Age: 18,
	}
	err := stuin.Store()
	if err != true{
		t.Fatalf("stroe测试错误%v", err)
	}
	t.Log("stroe测试正确")
}

func TestRestroe(t *testing.T)  {
	var stuin student // 接口
	stuin = &stu{}
	err, data := stuin.Restore()
	if err != true{
		t.Fatalf("restroe测试错误%v", err)
	}
	t.Logf("restroe测试正确反序列数据: %v", string(data))
}
