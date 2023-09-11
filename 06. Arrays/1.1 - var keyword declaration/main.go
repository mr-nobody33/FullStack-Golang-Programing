package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Jim"
	names[1] = "Kim"
	names[2] = "Alex"

	fmt.Println("Arrays Elements:")

	fmt.Println("Elements1:", names[0])
	fmt.Println("Elements2:", names[1])
	fmt.Println("Elements3:", names[2])
}
