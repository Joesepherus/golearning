package functions

import (
	"fmt"
	"sync"
)

func hello() (string, string) {
	return "hello", "world"
}

func functionWithMultipleReturns() {
	fmt.Println("function with multiple returns example")

	a, b := hello()
	fmt.Println(a, b)
}

func iife() {
	fmt.Println("iife example")

	func() {
		fmt.Println("Hello")
	}()
}

func closureBug() {
	fmt.Println("closure bug example")

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println("kokot", i)

		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

func closureFix() {
	fmt.Println("closure fix example")
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println("kokot", i)

		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func Run() {
	functionWithMultipleReturns()
	iife()
	closureBug()
	closureFix()
}
