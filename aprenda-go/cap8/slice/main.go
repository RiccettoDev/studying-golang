package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5} // aqui estamos declarando e atraibuildo valores a um array
	slice := []int{1, 2, 3, 4, 5}  // aqui estamos declarando e atraibuildo valores a um slice
	fmt.Println("Hello slice")
	fmt.Println(array)
	fmt.Println(slice)

	slice = append(slice, 6)

	fmt.Println("Adicionando um elemento extra para o slice")
	fmt.Println(slice)
	// a diferença é que o slice não é fixo e permite aumentar ou niminuir o número de elementos dentro dele.
}
