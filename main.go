package main

import "golearning/dependencyinjection2"

func main() {
	// fmt.Println(helloworld.Hello("Chris", ""))
	// repeated := iteration.Repeat("k", 10)
	// log.Print("repeated: ", repeated)
	// arrays.Slices()
	// arrays.DeepCopy()
	// arrays.DeepCopy2()
	// context3.Start()
	// goroutines.Start2()
	// dependencyinjection.Greet(os.Stdout, "Elodie")
	dependencyinjection2.Setup()
}
