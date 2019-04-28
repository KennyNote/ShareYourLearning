package main

func sum3(x, y int) int {
	s := x + y
	defer func() {
		s += 2
	}()
	return s
}

func sum4(x, y int) (s int) {
	s = x + y
	defer func() {
		s += 2
	}()
	return
}

func main() {
	s3 := sum3(1, 2)
	s4 := sum4(1, 2)
	println(s3) // 3
	println(s4) // 5
}
