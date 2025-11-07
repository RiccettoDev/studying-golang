package main

import (
	"fmt"
)

func main() {
	x := 11

	if x > 10 {
		fmt.Printf("X: %v é maior que 10\n", x)
	} else if x == 10 {
		fmt.Printf("X: %v é igual a 10\n", x)
	} else {
		fmt.Printf("X: %v é menor que 10\n", x)
	}
}
