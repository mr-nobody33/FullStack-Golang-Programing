package main

import "fmt"

func main() {
	a := 50
	b := 40

	r1 := a == b
	r2 := a != b
	r3 := a < b
	r4 := a <= b
	r5 := a > b
	r6 := a >= b

	fmt.Println("A = B :", r1)
	fmt.Println("A != B :", r2)
	fmt.Println("A < B :", r3)
	fmt.Println("A <= B :", r4)
	fmt.Println("A > B :", r5)
	fmt.Println("A >= B :", r6)
}
