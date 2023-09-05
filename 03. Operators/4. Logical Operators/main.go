package main

import "fmt"

func main() {
	var x, y, z = 100, 200, 300

	//Logical And ( && )
	fmt.Println(x < y && x > z)

	//Logical Or ( || )
	fmt.Println(x < y || x > z)

	//Logical Not ( ! )
	fmt.Println(!(x == y && x > z))
}
