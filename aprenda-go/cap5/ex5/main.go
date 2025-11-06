// Crie uma variável de tipo string utilizando uma raw string literal.
// Demonstre-a.

package main

import "fmt"

var raw string = `
	Isto é
		uma variável com
				raw string literal
`

func main() {
	fmt.Println(raw)
}
