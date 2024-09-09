package sets

import "fmt"

// Define a type alias for the set
type Set map[string]struct{}

// Function to create a new set
func NewSet() Set {
	return Set{}
}

// Function to add an element to the set
func (s Set) Add(element string) {
	s[element] = struct{}{}
}

// Function to remove an element from the set
func (s Set) Remove(element string) {
	delete(s, element)
}

// Function to check if an element exists in the set
func (s Set) Contains(element string) bool {
	_, exists := s[element]
	return exists
}

// Function to get the size of the set
func (s Set) Size() int {
	return len(s)
}

// Function to print all elements in the set
func (s Set) Print() {
	for element := range s {
		fmt.Println(element)
	}
}

func Run() {
	// Create a new set
	mySet := NewSet()

	// Add elements to the set
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("orange")

	// Print all elements in the set
	fmt.Println("Original set:")
	mySet.Print()

	// Check if an element exists
	fmt.Println("Contains 'banana':", mySet.Contains("banana")) // true
	fmt.Println("Contains 'grape':", mySet.Contains("grape"))   // false

	// Remove an element from the set
	mySet.Remove("orange")

	// Print updated set
	fmt.Println("Updated set:")
	mySet.Print()

	// Get the size of the set
	fmt.Println("Size of set:", mySet.Size())
}
