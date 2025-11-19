// Callback: passe uma função como argumento a outra função.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("A soma de 4 + 4 é?")
	fmt.Println(callbackFunction(soma(4, 4)))

	fmt.Println("--------")

	fmt.Println("A soma de 10 + 10 é?")
	result := parametroDeFunc(soma)
	fmt.Println(result(10, 10))
}

func soma(x, y int) int {
	return x + y
}

func callbackFunction(x int) (int, string) {
	return x, "é o resultado da soma"
}

func parametroDeFunc(z func(int, int) int) func(int, int) int {
	return z
}
