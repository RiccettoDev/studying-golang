// Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
// Por exemplo:
//     65
//         U+0041 'A'
//         U+0041 'A'
//         U+0041 'A'
//     66
//         U+0042 'B'
//         U+0042 'B'
//         U+0042 'B'
//     ...e por aí vai.

package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 122; i++ { // até 90 se for apenas as letras maiusculas e 122 caso queira maiusculas e minusculas
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("%#U\n", i)
		}
		fmt.Println("==================")
	}
}
