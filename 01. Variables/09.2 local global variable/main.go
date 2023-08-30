package main

import "fmt"

var globalVariable int = 100

func main() {
	// Local variable scope starts

	var localVariable int = 200

	fmt.Printf("The Value of Global Variable: %d\n", globalVariable)
	fmt.Printf("The Value of Local Variable: %d\n", localVariable)

	display()
}

func display() {
	fmt.Printf("The Value of Global Variable is: %d\n", globalVariable)
}
