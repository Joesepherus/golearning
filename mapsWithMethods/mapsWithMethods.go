package mapsWithMethods

import "fmt"

// Define a type alias for the dictionary (map)
type Dictionary map[string]int

// Function to create a new dictionary
func NewDictionary() Dictionary {
	return Dictionary{}
}

// Function to add or update a key-value pair
func (dictionary Dictionary) Add(key string, value int) {
	dictionary[key] = value
}

// Function to retrieve a value by key
func (dictionary Dictionary) Get(key string) (int, bool) {
	value, exists := dictionary[key]
	return value, exists
}

// Function to remove a key-value pair
func (dictionary Dictionary) Remove(key string) {
	delete(dictionary, key)
}

// Function to check if a key exists
func (dictionary Dictionary) Contains(key string) bool {
	_, exists := dictionary[key]
	return exists
}

// Function to get the size of the dictionary
func (dictionary Dictionary) Size() int {
	return len(dictionary)
}

// Function to print all key-value pairs
func (dictionary Dictionary) Print() {
	for key, value := range dictionary {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func Run() {
	// Create a new dictionary
	myDict := NewDictionary()

	// Add key-value pairs
	myDict.Add("apple", 5)
	myDict.Add("banana", 10)
	myDict.Add("orange", 3)

	// Print all key-value pairs
	fmt.Println("Original dictionary:")
	myDict.Print()

	// Access and modify values
	value, _ := myDict.Get("apple")
	fmt.Println("Apple count:", value)
	myDict.Add("grape", 7)
	myDict.Add("banana", 12) // Update value

	// Check for existence
	if value, exists := myDict.Get("pear"); exists {
		fmt.Println("Pear count:", value)
	} else {
		fmt.Println("Pear does not exist")
	}

	// Remove a key-value pair
	myDict.Remove("orange")

	// Print updated dictionary
	fmt.Println("Updated dictionary:")
	myDict.Print()

	// Get the size of the dictionary
	fmt.Println("Size of dictionary:", myDict.Size())
}
