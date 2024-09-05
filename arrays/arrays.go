package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func Slices() {
	mySlice := []string{"Pepper", "Elodie", "Ruth", "Pyewacket", "George"}

	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("mySlice[2:] = %v\n", mySlice[2:])
	fmt.Printf("mySlice[:2] = %v\n", mySlice[:2])
}

func DeepCopy() {
	x := [3]string{"Joseph", "Jack", "Dushan"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"

	z[1] = "Dennis" // the value at index 1 is now "Dennis" for only "z"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}

func DeepCopy2() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2]            // even though "b" len = 2, it points to the same the underlying array "a" points to
	b[1] = 1
	fmt.Println(b[1])
	fmt.Println(a[1])

	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
