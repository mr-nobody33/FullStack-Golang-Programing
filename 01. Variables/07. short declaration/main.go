package main

import "fmt"

func main() {
	// type of the variable is determined by the type of expression
	variable1 := 100
	variable2 := "Go Programing"
	variable3 := 135.94

	fmt.Printf("Variable1 value: %d\n", variable1)
	fmt.Printf("Variable1 type: %T\n", variable1)

	fmt.Printf("Variable2 value: %s\n", variable2)
	fmt.Printf("Variable2 type: %T\n", variable2)

	fmt.Printf("Variable3 value: %f\n", variable3)
	fmt.Printf("Variable3 type: %T\n", variable3)
}
