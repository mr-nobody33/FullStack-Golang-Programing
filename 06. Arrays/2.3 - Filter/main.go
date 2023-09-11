package main

import "fmt"

func main() {
	names := [5]string{"Kim", "Jim", "Bill", "Robert", "David"}

	fmt.Printf("Name: %v\n", names)

	// fmt.Println(names[:])
	fmt.Println(names[:3])  // Exibe apenas os 3 primeiros elementos.
	fmt.Println(names[2:])  // Exclui os 2 primeiros elementos.
	fmt.Println(names[1:4]) // Exclui o elemento 1 e mostra apenas até o 4º elemento.
}
