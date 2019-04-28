package main

import (
	"fmt"
)

func main() {

	type MyInt int

	var i interface{} = 12312

	i2 := i.(int)
	fmt.Println(i2)
	x1 := []int{
		1, 2, 3,
		4, 5, 6,
	}
	//x2 := []int{
	//	1, 2, 3,
	//	4, 5, 6
	//}
	x3 := []int{
		1, 2, 3,
		4, 5, 6}
	x4 := []int{
		1, 2, 3,
		4, 5, 6}
	x5 := []int{1, 2, 3, 4, 5, 6}
	x6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x1, x3, x4, x5, x6)

	ii := 1
	ii++
	fmt.Println(ii)
	flag := true
	if flag == false {
		println("123")
	} else {
		println("312")
	}

	var s []int
	fmt.Println(len(s), cap(s), s)
	s = append(s, 1)

	var m1 map[int]string = make(map[int]string, 5)
	m1[1] = "123"
	m1[2] = "234"
	todomap1(&m1)
	fmt.Println(&m1)

	var m2 map[int]string = make(map[int]string, 5)
	m2[1] = "123"
	m2[2] = "234"
	todomap2(m2)
	fmt.Println(m2)

	fmt.Println("+++++++++++++++++++")
	x := []string{"one", "two", "three"}
	for v, y := range x {
		fmt.Println(v, y)
	}

	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		fmt.Printf("w %s \n", s)
		go func() {
			fmt.Printf("n %s \n", s)
		}()
	}

}
func todomap1(m *map[int]string) {
	(*m)[2] = "432"
	(*m)[3] = "345"
}
func todomap2(m map[int]string) {
	m[2] = "432"
	m[3] = "345"
}
