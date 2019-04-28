package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int) {
	flag := true
	for i := 1; i < len(arr) && flag; i++ {
		flag = false
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				flag = true
			}
		}
	}
}

//二分查找
func BinarySearch(arr *[5]int, key int) (index int) {
	leftIndex := 0
	rightIndex := len(arr) - 1
	for leftIndex <= rightIndex {
		middle := (leftIndex + rightIndex) / 2
		if arr[middle] > key {
			rightIndex = middle - 1
		} else if arr[middle] < key {
			leftIndex = middle + 1
		} else {
			return middle
		}
	}
	return
}
func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println("main arr=", arr) //有序? 是有序的
	fmt.Println(BinarySearch(&arr, 80))
}
