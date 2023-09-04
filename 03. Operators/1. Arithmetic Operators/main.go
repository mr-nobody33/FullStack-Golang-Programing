package main

import "fmt"

func main() {
	a := 11
	b := 2

	addition := a + b
	subtraction := a - b
	multiplication := a * b
	division := a / b
	modulus := a % b

	fmt.Printf("Addition Result of a + b = %d\n", addition)
	fmt.Printf("Subtraction Result of a - b = %d\n", subtraction)
	fmt.Printf("Multiplication Result of a * b = %d\n", multiplication)
	fmt.Printf("Division Result of a / b = %d\n", division)
	fmt.Printf("Modulus Result of a %% b = %d\n", modulus)

	// Increment
	fmt.Println("Before Increment: ", a)
	a++
	fmt.Println("After Increment: ", a)

	// Decrement
	fmt.Println("Before Decrement: ", a)
	a--
	fmt.Println("After Decrement: ", a)
}
