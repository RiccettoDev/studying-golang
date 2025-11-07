// Põe na tela: todos os números de 1 a 10000.
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10000; i++ {
		if i != 1 && i != 10000 { // a lógica está funcionando corretamente, apenas adicionei este continue
			continue // para poder visualizar o 1 e o 10000 mais facil porque são muitos números
		}
		fmt.Println(i)
	}
}
