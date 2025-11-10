package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps range")

	qualquercoisa := map[int]string{
		123: "muito legal",
		456: "legal",
		789: "ruim",
	}

	fmt.Println(qualquercoisa)

	for key, value := range qualquercoisa { // aqui o indice, na verdade é um key, representan do a chave do valor
		fmt.Printf("Indice %v com o valor %v\n", key, value)
	}

	delete(qualquercoisa, 123) // assim deletamos um elemento de um maps

	fmt.Println(qualquercoisa)
}
