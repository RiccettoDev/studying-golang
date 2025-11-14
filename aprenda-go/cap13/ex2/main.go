// Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// Passe um valor do tipo slice of int como argumento para a função;
// Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// Passe um valor do tipo slice of int como argumento para a função.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(somaInt(2, 2, 2, 4))

	slice := []int{2, 2, 2}
	fmt.Println(somaInt(slice...))

	fmt.Println(somaSlice(slice))
}

func somaInt(x ...int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}

func somaSlice(x []int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}
