package binarySearch

import (
	"fmt"
	"log"
)

func Run() {
	log.Println("Binary Search")

	var arr = []int{1, 2, 4, 5, 7}
	var value int = 2

	res := binarySearch(arr, value)
	fmt.Printf("res %d\n\n\n", res)
}

func binarySearch(arr []int, value int) int {
	newArr := arr
	indexPlus := 0

	for len(newArr) > 0 {
		middle := len(newArr) / 2
		fmt.Println("middle", middle)
		fmt.Println("newArr[middle]", newArr, newArr[middle], value)

		if value == newArr[middle] {
			return indexPlus + middle
		}
		if len(newArr) == 1 {
			break
		}
		if value < newArr[middle] {
			newArr = newArr[:middle]
		} else {
			newArr = newArr[middle:]
			indexPlus += middle
		}
	}

	return -1
}
