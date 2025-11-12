package main

import (
	"fmt"
)

// método é uma função anexada a um tipo..
// e quando se anexa uma função a um tipo, ela se torna um método daquele tipo.

type pessoa struct {
	name string
	age  int
}

// estrutura de uma função:
// fun (receiver) identificador(parameters) (returnes) {code}

func (p pessoa) oibomdia() {
	fmt.Println(p.name, " diz bom dia!")
}

func main() {
	mauricio := pessoa{"Mauricio", 42}
	mauricio.oibomdia()
}

// um método adiciona funcionalidade especificamente à um tipo
// por exemplo se eu chamar o método oibomdia() sem estar atrelado à uma pessoa, ele não funciona
// portanto sua declaração funciona primeiro passando o valor do tipo e depois esta função método
// desta maneira: mauricio.oibomdia()
