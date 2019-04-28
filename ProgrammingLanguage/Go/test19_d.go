package main

import (
	"fmt"
	"sync"
)

func DataIn(a []int64, b []string, ch chan interface{}, wg *sync.WaitGroup) {
	for i := 0; i < 4; i++ {
		ch <- b[i]
		ch <- a[i]

	}
	close(ch)
	wg.Done()
}
func DataOut(ch chan interface{}, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	arra := []int64{1, 2, 3, 4}
	arrb := []string{"a", "b", "c", "d"}
	ch := make(chan interface{})
	wg.Add(2)
	go DataIn(arra, arrb, ch, &wg)
	go DataOut(ch, &wg)
	wg.Wait()
}
