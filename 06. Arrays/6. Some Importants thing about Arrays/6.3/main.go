package main

import "fmt"

func main() {
	// Creating an array whose size is determined by the number of elements present in it Using ellipsis.
	array := [...]string{"Go", "Java", "C#", "C++", "Perl"}

	fmt.Println("Elements of Array: ", array)

	// Length of the array is determine by Using len() function
	fmt.Println("Length of Array: ", len(array))
}
