package main

import "fmt"

func main() {

	// Variable declared and initialized without the explicit type
	var variable1 = 100                       // Int
	var variable2 = "Go Programming Language" // string
	var variable3 = 750.345                   // float

	// Display and print variables
	fmt.Printf("The value of Variable 1 : %d\n", variable1)
	fmt.Printf("The type of Variable 1 : %T\n", variable1)

	fmt.Printf("The value of Variable 2 : %s\n", variable2)
	fmt.Printf("The type of Variable 2 : %T\n", variable2)

	fmt.Printf("The value of Variable 3 : %f\n", variable3)
	fmt.Printf("The type of Variable 3 : %T\n", variable3)
}
