package main

import "fmt"

func main() {
	fmt.Println("From vars")

	var name string = "Mike"

	// Type infer
	var hello = "Bob"
	fmt.Println(name)
	fmt.Println(hello)

	// Printf takes in a format receiver
	fmt.Printf("%T\n", hello)
}