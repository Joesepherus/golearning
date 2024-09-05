package main

import (
	"fmt"
	"golearning/arrays"
	"golearning/helloworld"
	"golearning/iteration"
	"log"
)

func main() {
	fmt.Println(helloworld.Hello("Chris", ""))
	repeated := iteration.Repeat("k", 10)
	log.Print("repeated: ", repeated)
	arrays.Slices()
	arrays.DeepCopy()
	arrays.DeepCopy2()
}
