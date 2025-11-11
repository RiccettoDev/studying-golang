// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
// Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

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

	x := map[string]pessoa{
		"Pereira": {
			name:              "José",
			surname:           "Pereira",
			favoritesIceCream: []string{"chocolate", "milho verde"},
		},
		"Silva": {
			name:              "Maria",
			surname:           "Silva",
			favoritesIceCream: []string{"morango", "limão"},
		},
	}

	for sobrenome, pessoa := range x {
		fmt.Println("Sobrenome", sobrenome)
		fmt.Println("Nome", pessoa.name)
		fmt.Println("Sabores favoritos:")
		for i, sabor := range pessoa.favoritesIceCream {
			fmt.Printf("  %d. %s\n", i+1, sabor)
		}
		fmt.Println()
	}
}
