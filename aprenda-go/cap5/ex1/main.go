// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import (
	"fmt"
)

var number int = 10

func main() {
	fmt.Printf("Vamos exibir o número %v em Decimal, Binário e Hexadecimal", number)
	fmt.Println("=========================")
	fmt.Printf("Decimal: %d\n", number)
	fmt.Printf("Binário: %b\n", number)
	fmt.Printf("Hexadecimal: %#x\n", number)
}
