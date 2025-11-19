// Crie uma função que retorna uma função.
// Atribua a função retornada a uma variável.
// Chame a função retornada.

package main

import (
	"fmt"
)

func main() {
	x := retornaFuncao()
	x(10)
	fmt.Println(x(10)) // o retorno será 11
}

func retornaFuncao() func(int) int {
	return func(x int) int {
		return x + 1
	}
}
