package main

import (
	"fmt"
)

type People8_2 interface {
	Speak(string) string
}

type Stduent8_2 struct{}

func (stu Stduent8_2) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo1 People8_2 = Stduent8_2{}
	var peo2 People8_2 = &Stduent8_2{}
	think := "bitch"
	//如果一个medthod的接收者定义为值类型 那么在调用的时候 传入值或者指针都可以  如果定义的是指针类型 那么只能为指针类型。也就是说在本例的代码中 如果传入&u 那么就不会有错误。
	fmt.Println(peo1.Speak(think))
	fmt.Println(peo2.Speak(think))
}
