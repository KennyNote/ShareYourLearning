package main

import (
	"fmt"
	"time"
)

func main() {
	ChannelFunc()
}
func ChannelFunc() {
	zimu := make(chan int, 1)
	suzi := make(chan int, 1)
	zimu <- 0
	// zimu
	go func() {
		for i := 65; i <= 90; i++ {
			<-zimu
			fmt.Printf("%v", string(rune(i)))
			suzi <- i
		}
		return
	}()

	go func() {
		for i := 1; i <= 26; i++ {
			<-suzi
			fmt.Printf("%v", i)
			zimu <- i
		}
		return
	}()

	time.Sleep(1 * time.Second)
	fmt.Println()
}
