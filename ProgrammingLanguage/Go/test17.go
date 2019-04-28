package main

import "fmt"

type Test struct {
	Name string
}

var list map[string]*Test

var list2 map[string]string

var list3 map[string]Test

func main() {

	list = make(map[string]*Test)
	name := Test{"xiaoming"}
	list["name"] = &name
	list["name"].Name = "Hello"
	fmt.Println(list["name"])

	list2 = make(map[string]string)
	list2["name"] = "qwe"
	fmt.Println(list2["name"])

	list3 = make(map[string]Test)
	name3 := Test{"xiaoming"}
	list3["name"] = name3
	//list3["name"].Name = "Hello"
	fmt.Println(list["name"])

}
