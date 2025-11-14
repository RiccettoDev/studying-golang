// Atribua uma função a uma variável.
// Chame essa função.

package main

import (
	"fmt"
)

var x string

func main() {
	x = helloWord()

	fmt.Println(x)

	y := func() {
		fmt.Println("Atribuindo de outra forma")
	}

	y()
}

func helloWord() string {
	return "Hello World!"
}
