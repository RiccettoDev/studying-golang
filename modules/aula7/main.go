package main // pacote principal

import (
	"fmt"
)

// Condicionais
// IF / ELSE
// SE / SENÃO

func main() {
	valor := 2
	// if (condição) { <ação> }
	if valor == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor é diferente de 1")
	}

	x := 10
	if x > 5 {
		fmt.Println("x é maior que 5")
	} else if x < 5 {
		fmt.Println("x é menor que 5")
	} else {
		fmt.Println("x é igual a 5")
	}
}
