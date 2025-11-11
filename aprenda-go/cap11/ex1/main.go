// Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// Nome
// Sobrenome
// Sabores favoritos de sorvete
// Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

package main

import (
	"fmt"
)

type pessoa struct {
	name              string
	surname           string
	favoritesIceCream []string
}

func main() {
	pessoa1 := pessoa{
		name:              "José",
		surname:           "Pereira",
		favoritesIceCream: []string{"cholate", "milho verde"},
	}

	pessoa2 := pessoa{
		name:              "Maria",
		surname:           "Silva",
		favoritesIceCream: []string{"morango", "limão"},
	}

	fmt.Printf("- %v %v tem como sabores de sorvete favorito: \n", pessoa1.name, pessoa1.surname)
	for i, v := range pessoa1.favoritesIceCream {
		fmt.Printf("%v - %v\n", i+1, v)
	}

	fmt.Println("")
	fmt.Println("------------")
	fmt.Println("")

	fmt.Printf("- %v %v tem como sabores de sorvete favorito: \n", pessoa2.name, pessoa2.surname)
	for i, v := range pessoa2.favoritesIceCream {
		fmt.Printf("%v - %v\n", i+1, v)
	}
}
