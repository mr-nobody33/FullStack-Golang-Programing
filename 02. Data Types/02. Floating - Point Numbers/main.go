package main

import "fmt"

func main() {
	// Define floating-point numbers
	a := 35.78
	b := 46.26

	c := b - a

	fmt.Printf("A = %f\n", a)
	fmt.Printf("B = %f\n", b)

	fmt.Printf("%f - %f = %f\n", a, b, c)
	fmt.Printf("Type of C: %T\n", c)
}
