package main // pacote principal

import (
	"fmt"
)

// Switch Case

func main() {
	posicao := 5

	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	default:
		fmt.Println("Outro lugar")
	}

}
