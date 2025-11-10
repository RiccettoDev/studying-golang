// Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
// "Nome", "Sobrenome", "Hobby favorito"
// Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		{"Joao", "Pereira", "Programar"},
		{"Maria", "Silva", "Desenhar"},
		{"Lucas", "Almeida", "Cantar"},
	}

	for _, value := range ss {
		fmt.Printf("Nome: %v, Sobrenome: %v, Hobby: %v\n", value[0], value[1], value[2])
	}
}
