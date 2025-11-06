package main

import (
	"fmt"
)

func main() {
	a := 100

	fmt.Printf("Valores em diferentes formatos:\n")
	fmt.Printf("\n")
	fmt.Printf("Valor de 100 em decimal: %d\n", a)
	fmt.Printf("Valor de 100 em binário: t%b\n", a)
	fmt.Printf("Valor de 100 em hexadecimal: t%#x\n", a)
}
