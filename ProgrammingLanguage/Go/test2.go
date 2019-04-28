package main

import "fmt"

func main() {
	pase_student()
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	m2 := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 错误写法
	for _, stu := range stus {
		m[stu.Name] = &stu
		m2[stu.Name] = stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("=============")
	for k, v := range m2 {
		fmt.Println(k, "=>", v.Name)
	}

	fmt.Println("=============")
	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}
