// Usando uma literal composta:
// Crie um array que suporte 5 valores to tipo int
// Atribua valores aos seus índices
// Utilize range e demonstre os valores do array.
// Utilizando format printing, demonstre o tipo do array.

package main

import (
	"fmt"
)

func main() {

	var array [5]int

	for i := 0; i < len(array); i++ {
		array[i] = i * 2
	}

	fmt.Println(array)

	for i, v := range array {
		fmt.Printf("No indice %v, tem o valor %v, no tipo %T\n", i, v, v)
	}
}
