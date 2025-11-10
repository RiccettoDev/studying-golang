package main

import (
	"fmt"
)

func main() {

	slice := []string{"banana", "maça", "cenoura"}

	fmt.Println(slice)

	slice = append(slice, "uva")

	fmt.Println(slice)

	fmt.Println("Substituindo a uva por abacare e fazendo um for range no slice")

	slice[3] = "abacate" // aqui substuimos a uva na posição 3 pelo abacate

	for i, valor := range slice {
		fmt.Printf("No indice %v temos o valor %v\n", i, valor)
	}

	// o simpulo de = vai fazer atribuição de um valor em um campo existente do slice
	// mas apara criar um novo campk é necessário utilizar o append
}
