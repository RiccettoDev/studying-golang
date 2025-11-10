// Crie um map com key tipo string e value tipo []string.
// Key deve conter nomes no formato sobrenome_nome
// Value deve conter os hobbies favoritos da pessoa
// Demonstre todos esses valores e seus índices.

package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"Joao Pereira":   {"correr, malhar"},
		"Maria Silva":    {"desenhar, cantar"},
		"Carlos Antonio": {"ler, estudar"},
	}

	for i, v := range x {
		fmt.Printf("%v: %v\n", i, v)
	}

	fmt.Println("")
	fmt.Println("================")
	fmt.Println("Deletar um item do maps")
	fmt.Println("================")
	fmt.Println("")

	delete(x, "Carlos Antonio")

	for i, v := range x {
		fmt.Printf("%v: %v\n", i, v)
	}
}
