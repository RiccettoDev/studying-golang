package main

import (
	"fmt"
)

// no go, para loop temos apenas o FOR

// aqui temos um for padrão
func forPadrao(a int) {
	for i := 0; i < a; i++ {
		println("Interação: ", i)
	}
}

// aqui temos um for range
func forRange() {
	numeros := []int{1, 2, 3, 4, 5}
	for k, v := range numeros { // k de chave, v de valor.  E podemos esconder qualquer um deles substituindo por _
		fmt.Printf("Posição %v tem o valor %v \n", k, v)
	}
}

// aqui um for tipo while
func forWhile() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}

// for de loop infinito
func forInfinit() {
	for {
		println("Hello")
	}
}

func main() {
	forPadrao(10)
	forRange()
	forWhile()
	forInfinit()
}
