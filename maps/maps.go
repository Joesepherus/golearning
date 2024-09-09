package maps

import "fmt"

func Run() {
	// Create a dictionary (map) with string keys and integer values
	myDict := map[string]int{
		"apple":  5,
		"banana": 10,
		"orange": 3,
	}

	// Accessing values
	fmt.Println("Apples:", myDict["apple"]) // 5

	// Adding a new key-value pair
	myDict["grape"] = 7
	fmt.Println("Grapes:", myDict["grape"]) // 7

	// Modifying a value
	myDict["banana"] = 12
	fmt.Println("Bananas:", myDict["banana"]) // 12

	// Checking if a key exists
	value, exists := myDict["pear"]
	if exists {
		fmt.Println("Pears:", value)
	} else {
		fmt.Println("Pears do not exist")
	}

	// Removing a key-value pair
	delete(myDict, "orange")

	myDict["apple"] = 99

	// Looping through a dictionary (map)
	fmt.Println("Updated dictionary:")
	for key, value := range myDict {
		fmt.Printf("%s: %d\n", key, value)
	}
}
