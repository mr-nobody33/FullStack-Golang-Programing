package main

import "fmt"

func main() {
	var myMap = make(map[string]float64)

	myMap["Kim"] = 80.55
	myMap["Jim"] = 90.33
	myMap["Robert"] = 75.60

	fmt.Println("MyMap: ", myMap)
}
