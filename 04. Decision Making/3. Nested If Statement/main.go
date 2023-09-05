package main

import "fmt"

func main() {
	var a int = 300
	var b int = 500

	if a == 300 {
		if b == 500 {
			fmt.Println("Value A is 300 and Value B is 500!")
		}
	}

	fmt.Printf("A Value: %d\n", a)
	fmt.Printf("B Value: %d\n", b)
}
