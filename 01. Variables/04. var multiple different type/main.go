package main

import "fmt"

func main() {
	//declare multiple variables for different type in the single diclaration
	var variable1, variable2, variable3 = 2, "Golang", 67.56

	fmt.Printf("Variable1 value: %d\n", variable1)
	fmt.Printf("Variable1 type: %T\n", variable1)

	fmt.Printf("Variable1 value: %s\n", variable2)
	fmt.Printf("Variable1 type: %T\n", variable2)

	fmt.Printf("Variable1 value: %f\n", variable3)
	fmt.Printf("Variable1 type: %T\n", variable3)
}
