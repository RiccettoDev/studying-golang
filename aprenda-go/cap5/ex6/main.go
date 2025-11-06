// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// Demonstre estes valores.

package main

import (
	"fmt"
)

const (
	anoAtual int = 2025 + iota
	anoB
	anoC
	anoD
	anoE
)

func main() {
	fmt.Printf("Este é o ano atual: %v\n", anoAtual)
	fmt.Printf("Este é o próximo ano atual: %v\n", anoB)
	fmt.Printf("Este é o próximo ano atual: %v\n", anoC)
	fmt.Printf("Este é o próximo ano atual: %v\n", anoD)
	fmt.Printf("Este é o próximo ano atual: %v\n", anoE)

}
