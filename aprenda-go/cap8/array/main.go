package main

import (
	"fmt"
)

var x [5]int // aqui estamos declarando um array de 5 números inteiros

// Arrays possuiem números de elementos fixos e não é possível aumentar

func main() {
	fmt.Println("Hello Array!")
	fmt.Println(x)
	x[0] = 5
	x[1] = 10
	fmt.Println(x)
	fmt.Printf("O tipo do array é: %T \n", x)
	fmt.Printf("O tamnho do array é: ")
	fmt.Println(len(x))

	// saídas no terminal apos a execução dos comandos:

	// Hello Array!
	// [0 0 0 0 0]
	// [5 10 0 0 0]
	// O tipo do array é: [5]int
	// O tamnho do array é: 5
}
