package main

import (
	"fmt"
	"golearning/helloworld"
	"golearning/iteration"
	"log"
)

func main() {
	fmt.Println(helloworld.Hello("Chris", ""))
	repeated := iteration.Repeat("k")
	log.Print("repeated: ", repeated)
}
