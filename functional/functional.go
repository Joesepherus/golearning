package functional

import (
	"fmt"
)

// A generic function that takes a slice and a function to process each element
func processTasks(tasks []int, processor func(int) int) int {
	result := 0
	for _, task := range tasks {
		result += processor(task)
	}
	return result
}

func double(task int) int {
	return task * 2
}

// Function to process a single task (identity function here)
func identity(task int) int {
	return task
}

func Run() {
	// Define the tasks (list of integers)
	tasks := []int{1, 2, 3, 4, 5}

	// Process the tasks using a higher-order function
	// result := processTasks(tasks, identity)
	summResult := processTasks(tasks, identity)
	doubleResult := processTasks(tasks, double)

	// Print the result
	fmt.Println("The sum of tasks is:", summResult)
	fmt.Println("The sum of tasks is:", doubleResult)
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(n int) bool {
	return n%2 == 1
}

func filter(nums []int, condition func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if condition(n) {
			result = append(result, n)
		}
	}
	return result
}

// Functions as First-Class Citizens
func Run2() {
	nums := []int{1, 2, 3, 4, 5}
	evens := filter(nums, isEven)
	fmt.Println(evens) // Output: [2, 4]
	odds := filter(nums, isOdd)
	fmt.Println(odds) // Output: [2, 4]
}

func mapFunc(nums []int, f func(int) int) []int {
	var result []int
	for _, n := range nums {
		result = append(result, f(n))
	}
	return result
}

// This is a pure function
func square(n int) int {
	return n * n
}

// Higher-Order Functions
func Run3() {
	nums := []int{1, 2, 3}
	squared := mapFunc(nums, square)
	fmt.Println(squared) // Output: [1, 4, 9]
}

// Imperative way
func imperativeAddOne(nums []int) []int {
	var result []int
	for _, n := range nums {
		result = append(result, n+1)
	}
	return result
}

// Declarative way using a map function
func declarativeAddOne(nums []int) []int {
	return mapFunc(nums, func(n int) int {
		return n + 1
	})
}

// Declarative vs. Imperative Thinking
func Run4() {
	nums := []int{1, 2, 3}
	addedOne := imperativeAddOne(nums)
	fmt.Println(addedOne)
	addedOneDeclarative := declarativeAddOne(nums)
	fmt.Println(addedOneDeclarative)
}

// Function Composition
func Run5() {
	nums := []int{1, 2, 3, 4, 5}
	evens := filter(nums, isEven)
	doubledEvens := mapFunc(evens, double)
	fmt.Println(doubledEvens) // Output: [4, 8]
}
