package main

import "fmt"

func main() {
	var emptyMap = map[int]string{}

	fmt.Println("Empty Map: ", emptyMap)

	var myMap = map[int]string{
		1: "Kim",
		2: "Jim",
		3: "Robert",
		4: "Jim",
		5: "Edvard",
	}

	fmt.Println("My Map: ", myMap)
}
