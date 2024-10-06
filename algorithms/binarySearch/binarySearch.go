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

	res = binarySearchProper(arr, value)
	fmt.Println("res", res)
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

func binarySearchProper(arr []int, target int) int {
	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2 // To avoid overflow
		fmt.Println("mid", mid)
		fmt.Println("start, end ", start, end)
		fmt.Println("arr arr[mid], value", arr[start:end+1], arr[mid], target)

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1 // Target not found
}
