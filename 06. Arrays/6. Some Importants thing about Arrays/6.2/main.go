package main

import "fmt"

func main() {
	// Creating array Using shorthand declaration
	array1 := [3]int{1, 2, 3}
	array2 := [...]int{1, 2, 3, 4, 5}

	// Finding the length of the array using len function
	fmt.Println("Length of Array1: ", len(array1))
	fmt.Println("Length of Array2: ", len(array2))
}
