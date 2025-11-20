package main

import (
	"encoding/json"
	"fmt"
)

type pessoa struct { // seguindo a mesma ógica de termos de deixar métodos e funções como maiúscula para exportar
	Name  string // temos de deixar os atributos com a primeira maiúscula, caso contrário não exporta para json
	Age   int    // caso algum dos atributos não esteja em maiúsculo, ele não será exportado.
	email string // por exemplo o email está em minúscula e nunca sera exportado para o json
}

func main() {

	// podemos declarar assim:
	jose := pessoa{"José", 30, "jose@gmail.com"}
	carlos := pessoa{"Carlos", 35, "carlos@gmail.com"}

	// ou podemos declarar assim também:
	maria := pessoa{Name: "Maria", Age: 25, email: "maria@gmail.com"}
	ana := pessoa{Name: "Ana", Age: 31, email: "ana@gmail.com"}

	j, err := json.Marshal(jose)
	if err != nil {
		fmt.Println(err)
	}

	c, err := json.Marshal(carlos)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(maria)
	if err != nil {
		fmt.Println(err)
	}

	a, err := json.Marshal(ana)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(j) // desta maneira teremos um slice of bytes, onde o resultado é este:
	fmt.Println(c) // [123 34 78 97 109 101 34 58 34 67 97 114 108 111 115 34 44 34 65 103 101 34 58 51 53 125]
	fmt.Println(m)
	fmt.Println(a)

	fmt.Println(string(j)) // com o string() mostramos o json
	fmt.Println(string(c))
	fmt.Println(string(m))
	fmt.Println(string(a))
}

// resultados do terminal:
// [123 34 78 97 109 101 34 58 34 74 111 115 195 169 34 44 34 65 103 101 34 58 51 48 125]
// [123 34 78 97 109 101 34 58 34 67 97 114 108 111 115 34 44 34 65 103 101 34 58 51 53 125]
// [123 34 78 97 109 101 34 58 34 77 97 114 105 97 34 44 34 65 103 101 34 58 50 53 125]
// [123 34 78 97 109 101 34 58 34 65 110 97 34 44 34 65 103 101 34 58 51 49 125]
// {"Name":"José","Age":30}
// {"Name":"Carlos","Age":35}
// {"Name":"Maria","Age":25}
// {"Name":"Ana","Age":31}
