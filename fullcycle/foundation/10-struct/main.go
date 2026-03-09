package main

import (
	"fmt"
)

// uma struct lembra muito uma classe

type Endereço struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	name     string
	age      int
	active   bool
	Endereço          // pode ser usacomo como uma composição ou como um tipo
	Adress   Endereço // aqui sendo usado como um tipo
}

func (c *Client) Desativar() { // aqui criamos um método para o struct Client (receiver por ponteiro)
	c.active = false
}

func main() {
	joao := Client{name: "Joao", age: 32, active: true}
	joao.active = false
	fmt.Println(joao)

	joao.Logradouro = "Rua das margaridas"
	joao.Numero = 100
	joao.Cidade = "Santana de Parnaiba"
	joao.Estado = "Pernambuco"
	joao.Adress.Cidade = "Rua 1"
	joao.active = true
	fmt.Println(joao)

	joao.Desativar()

	fmt.Println(joao)

}
