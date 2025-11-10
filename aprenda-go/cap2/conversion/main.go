package main

import (
	"fmt"
)

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	// fmt.Printf("%v\n", x)

	x = int(b) // colocando o tipo que vc deseja converter o o valor dentro dele é possivel fazer a conversão
	// por exemplo aqui colocando b dentro de int() convertemos o tipo de hotdog para int...
	fmt.Printf("%v\n", x)

	//fmt.Println(x == b) // aqui propositalmente estamos deixando assim, pois a ide informa o erro de que os tipos são diferentes nesta operação
}
