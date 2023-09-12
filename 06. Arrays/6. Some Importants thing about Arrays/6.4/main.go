package main

import "fmt"

func main() {
	// Creating an array whose size is represented by the ellipsis
	array := [...]int{10, 20, 30, 40, 50, 60}

	// Iterate array using for loop
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\n", array[i])
	}
}
