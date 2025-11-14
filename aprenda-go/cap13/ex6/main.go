// Crie e utilize uma função anônima.

package main

import (
	"fmt"
)

func main() {

	x := func(y, z int) int {
		result := y + z
		return result
	}

	fmt.Println("Função anônima para soma")
	fmt.Println(x(2, 2))
}
