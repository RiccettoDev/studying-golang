// Crie um tipo struct "pessoa" que contenha os campos:
// nome
// sobrenome
// idade
// Crie um método para "pessoa" que demonstre o nome completo e a idade;
// Crie um valor de tipo "pessoa";
// Utilize o método criado para demonstrar esse valor.

package main

import (
	"fmt"
)

type pessoa struct {
	name    string
	surname string
	age     int
}

func (p pessoa) apresentacao() {
	fmt.Printf("Nome: %s\nSobrenome: %s\nIdade: %d\n", p.name, p.surname, p.age)
}

func main() {
	fmt.Println("")
	jose := pessoa{
		name:    "José",
		surname: "Pereira",
		age:     55,
	}

	jose.apresentacao()
}
