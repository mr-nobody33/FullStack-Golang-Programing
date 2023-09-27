package main

import "fmt"

func main() {
	// Creating empty map
	var emptyMap = map[int]string{}

	fmt.Println("Empty Map: ", emptyMap)

	// Creating and initializing a simple map
	var myMap = map[int]string{
		1: "Kim",
		2: "Jim",
		3: "Robert",
		4: "Jim",
		5: "Edvard",
	}

	fmt.Println("My Map: ", myMap)
}
