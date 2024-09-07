package goroutines

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func Start() {
	go printMessage("Hello")
	go printMessage("World")

	// Wait for goroutines to finish
	time.Sleep(1 * time.Second)
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func Start2() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
}
