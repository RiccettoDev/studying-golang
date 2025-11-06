// iota é uma palavra reservada usada dentro de um bloco const,
// que gera valores inteiros sequenciais automaticamente, começando em 0.

package main

import "fmt"

func main() {
	const (
		A = iota // 0
		B        // 1
		C        // 2
	)
	fmt.Println(A, B, C)
}
