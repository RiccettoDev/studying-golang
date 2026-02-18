package main

import (
	"errors"
	"fmt"
)

func main() {
	println(sum(1, 3))
	hello("Edu")
	fmt.Println(sumVerdade(1, 3))

	valor, err := validation(20, 20) // itilização da função validation
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
}

func hello(x string) { // funções podem ter retorno ou não...  neste caso aqui não tem retorno
	fmt.Printf("Hello %v!!!\n", x)
}

func sum(a int, b int) int { // mas neste caso tem, então preciso especificar o tipo do retorno, exatamente como fiz após os parenteses
	return a + b
}

func sumVerdade(a int, b int) (int, bool) { // podemos tambem ter mais de um retorno e de tipos diferentes

	soma := a + b

	return soma, a == b
}

func validation(a int, b int) (int, error) { // como em gol não existe um try catch existem outras maneiras de fazermos validações
	if a+b >= 50 {
		return 0, errors.New("A soma é igual ou maior que 50")
	}

	return a + b, nil // nil em Go é o valor zero (ou valor padrão) para tipos que podem ser "vazios" ou "não inicializados"... usamos aqui pq não temos erro, sendo que o retorno da função esperava por um erro.
	// podemos entender o nil como um null
}
