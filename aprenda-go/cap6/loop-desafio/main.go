// Desafio surpresa!
// Format printing:
// Decimal       %d
// Hexadecimal   %#x
// Unicode       %#U
// Tab           \t
// Linha nova    \n
// Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.

package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Valor %v em Decimal é %v \n", i, i)
		fmt.Printf("Valor %v em Hexadecimal é %#x \n", i, i)
		fmt.Printf("Valor %v em Unicode é %#U \n", i, i)
		fmt.Println("==============================================")
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("Valor %v somente a string %v \n", i, string(i))
		fmt.Println("==============================================")
	}
}
