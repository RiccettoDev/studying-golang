package main

import (
	"fmt"
)

var y = 30

func main() {
	fmt.Println("Hello Golang")

	x := 22

	fmt.Printf("X tem o valor: %v e o tipo %T\n", x, x) // printf é o format printer...
	// e desta forma podemos formatar para exibir valor e tipo da variável

	fmt.Printf("Y tem o valor: %v e o tipo %T\n", y, y) // printf é o format printer...

}

// A mormota := operador curto de declaração funciona apenas dentro de um escopo curto (dentro de um função), ela não possui escopo global dentro de um pacote
// patra uma declaração de escopo global em um pacote o correto deria var y = 30
