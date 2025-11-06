package main

import "fmt"

func main() {
	a := 2      // 010 em binário
	b := a >> 1 // desloca 1 bit para a direita
	c := a << 1 // desloca 1 bit para a esquerda

	fmt.Printf("a = %03b (%d)\n", a, a)
	fmt.Printf("b = %03b (%d)\n", b, b)
	fmt.Printf("c = %03b (%d)\n", c, c)
}

// tendo este resultado como resposta
// a = 010 (2)
// b = 001 (1)
// c = 100 (4)
