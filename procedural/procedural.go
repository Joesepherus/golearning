package procedural

import (
	"fmt"
)

// Function to process tasks (summing numbers)
func processTasks(tasks []int) int {
	sum := 0
	for _, task := range tasks {
		sum += task
	}
	return sum
}

func Run() {
	// Define the tasks (list of integers)
	tasks := []int{1, 2, 3, 4, 5}

	// Process the tasks and get the result
	result := processTasks(tasks)

	// Print the result
	fmt.Println("The sum of tasks is:", result)
}
