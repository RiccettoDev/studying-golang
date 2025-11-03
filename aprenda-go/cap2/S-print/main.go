package main

import (
	"fmt"
)

func main() {

	a := "Olá"
	b := "Mundo!"

	x := fmt.Sprint(a, " ", b) // o Sprint é uma concatenação de strings, onde estou juntando uma única string atraves da
	// concatenação de a + " " + b que gera a string: Olá Mundo!

	fmt.Println(x)
}
