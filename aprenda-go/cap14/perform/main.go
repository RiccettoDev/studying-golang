// ponteiro serve para performar
// em GO ao passar valor para uma variável custa performa
// ao invés de passar o valor utilizamos ponteiros para indicar à variável de onde ela deve usar o valor
// assim otimizamos a performance pois não estamos passando o valor e sim apontando onde está o valor à ser utilizado

// Um ponteiro é uma variável que guarda o endereço na memória de outra variável.
// Ou seja, em vez de trabalhar com uma cópia do valor, você trabalha diretamente com o valor ORIGINAL, permitindo modificá-lo.

package main

import (
	"fmt"
)

func main() {
	x := 0

	semPonteiro(x) // incrementa só a cópia
	fmt.Println("Depois de semPonteiro:", x)

	comPonteiro(&x) // incrementa o valor original
	fmt.Println("Depois de comPonteiro:", x)
}

func semPonteiro(x int) {
	x++
	fmt.Println("Dentro de semPonteiro:", x)
}

func comPonteiro(x *int) {
	*x++ // altera o valor no endereço
	fmt.Println("Dentro de comPonteiro:", *x)
}
