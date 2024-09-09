package controlflows

import "fmt"

func simpleLoop() {
	fmt.Println("simple loop example")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func loopOverArray() {
	fmt.Println("loop over array example")

	arr := []int{1, 2, 3, 4, 5}
	for i, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	fmt.Println("loop over array example 2")

	for i, v := range []string{"Rick", "Morty", "Beth", "Summer", "Jerry"} {
		fmt.Printf("%v at index %d\n", v, i)
	}
}

func whileLoop() {
	fmt.Println("while loop example")

	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}
}

func ifelseif() {
	fmt.Println("if else if example")

	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}
}

func switchExample() {
	fmt.Println("switch example")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}
}

// Acts as if-else
func switchWithoutExpression() {
	fmt.Println("switch without expression example")

	x := 10
	switch {
	case x > 0:
		fmt.Println("x is positive")
	case x < 0:
		fmt.Println("x is negative")
	default:
		fmt.Println("x is zero")
	}

}

func switchWithMultipleCases() {
	fmt.Println("switch with multiple cases example")

	number := 2
	switch number {
	case 1, 3, 5:
		fmt.Println("Odd number")
	case 2, 4, 6:
		fmt.Println("Even number")
	default:
		fmt.Println("Other number")
	}

}

func Run() {
	simpleLoop()
	loopOverArray()
	whileLoop()
	ifelseif()
	switchExample()
	switchWithoutExpression()
	switchWithMultipleCases()
}
