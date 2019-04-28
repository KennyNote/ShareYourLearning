package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(987654321))
	fmt.Println(reverseString("qwertyuiop"))
}
func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return
}
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
