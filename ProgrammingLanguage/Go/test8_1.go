package main

import (
	"fmt"
)

type People8_1 interface {
	Speak(string) string
}

type Stduent8_1 struct{}

func (stu *Stduent8_1) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People8_1 = &Stduent8_1{}
	think := "bitch"
	//如果一个medthod的接收者定义为值类型 那么在调用的时候 传入值或者指针都可以  如果定义的是指针类型 那么只能为指针类型。也就是说在本例的代码中 如果传入&u 那么就不会有错误。
	fmt.Println(peo.Speak(think))
}
