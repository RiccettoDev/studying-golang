// Crie uma função que retorne um int
// Crie outra função que retorne um int e uma string
// Chame as duas funções
// Demonstre seus resultados

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Uma função que retorna um INT")
	fmt.Println(returnInt(5))
	fmt.Println("Uma função que retorna um INT e uma STRING")
	fmt.Println(returnIntAndString(10))
}

func returnInt(x int) int {
	return x
}

func returnIntAndString(y int) (int, string) {
	return y, "E aqui o resultado string"
}
