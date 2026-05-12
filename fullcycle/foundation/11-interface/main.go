package main

import (
	"fmt"
)

// ============================================
// CONCEITO: Interface em Go
// ============================================
// Interface define um CONTRATO - um conjunto de métodos que um tipo deve implementar
// Em Go, a implementação é IMPLÍCITA - não precisa declarar "implements"
// Se um tipo tem todos os métodos da interface, ele automaticamente a implementa

// ============================================
// Exemplo 1: Interface básica
// ============================================

type Client struct {
	Name   string
	Age    int
	Active bool
}

type Company struct {
	Name   string
	CNPJ   string
	Active bool
}

// Métodos para Client
func (c *Client) Desativar() {
	c.Active = false
	fmt.Printf("Cliente %s desativado\n", c.Name)
}

func (c *Client) Ativar() {
	c.Active = true
	fmt.Printf("Cliente %s ativado\n", c.Name)
}

// Métodos para Company
func (c *Company) Desativar() {
	c.Active = false
	fmt.Printf("Empresa %s desativada\n", c.Name)
}

func (c *Company) Ativar() {
	c.Active = true
	fmt.Printf("Empresa %s ativada\n", c.Name)
}

// Interface que define o contrato
// Qualquer tipo que tenha os métodos Desativar() e Ativar() implementa essa interface
type Pessoa interface {
	Desativar()
	Ativar()
}

// Função que aceita qualquer tipo que implemente a interface Pessoa
func DesativarPessoa(p Pessoa) {
	p.Desativar()
}

// ============================================
// Exemplo 2: Interface vazia (interface{} ou any)
// ============================================
// A interface vazia pode receber QUALQUER tipo

func PrintQualquerCoisa(i interface{}) {
	fmt.Printf("Tipo: %T, Valor: %v\n", i, i)
}

// ============================================
// Exemplo 3: Type Assertion
// ============================================
// Usado para extrair o valor concreto de uma interface

func VerificarTipo(i interface{}) {
	// Type assertion com verificação
	if v, ok := i.(string); ok {
		fmt.Printf("É uma string: %s\n", v)
	} else if v, ok := i.(int); ok {
		fmt.Printf("É um int: %d\n", v)
	} else {
		fmt.Println("Tipo desconhecido")
	}
}

// ============================================
// Exemplo 4: Type Switch
// ============================================

func IdentificarTipo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("String de tamanho %d: %s\n", len(v), v)
	case int:
		fmt.Printf("Inteiro: %d\n", v)
	case Client:
		fmt.Printf("Cliente: %s, idade %d\n", v.Name, v.Age)
	case *Client:
		fmt.Printf("Ponteiro para Cliente: %s\n", v.Name)
	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
}

// ============================================
// Exemplo 5: Interface com múltiplos métodos
// ============================================

type Animal interface {
	Falar() string
	Mover() string
}

type Cachorro struct {
	Nome string
}

type Passaro struct {
	Nome string
}

func (c Cachorro) Falar() string {
	return "Au au!"
}

func (c Cachorro) Mover() string {
	return "Correndo"
}

func (p Passaro) Falar() string {
	return "Piu piu!"
}

func (p Passaro) Mover() string {
	return "Voando"
}

func FazerAnimalAgir(a Animal) {
	fmt.Printf("%s - %s\n", a.Falar(), a.Mover())
}

func main() {
	fmt.Println("=== Exemplo 1: Interface básica ===")
	joao := Client{Name: "João", Age: 32, Active: true}
	empresa := Company{Name: "Tech Corp", CNPJ: "12345678000190", Active: true}

	// Ambos implementam a interface Pessoa
	DesativarPessoa(&joao)
	DesativarPessoa(&empresa)

	fmt.Println("\n=== Exemplo 2: Interface vazia ===")
	PrintQualquerCoisa("Hello")
	PrintQualquerCoisa(42)
	PrintQualquerCoisa(joao)
	PrintQualquerCoisa(true)

	fmt.Println("\n=== Exemplo 3: Type Assertion ===")
	VerificarTipo("Golang")
	VerificarTipo(100)
	VerificarTipo(3.14)

	fmt.Println("\n=== Exemplo 4: Type Switch ===")
	IdentificarTipo("Interface em Go")
	IdentificarTipo(2024)
	IdentificarTipo(joao)
	IdentificarTipo(&joao)

	fmt.Println("\n=== Exemplo 5: Polimorfismo com Interface ===")
	dog := Cachorro{Nome: "Rex"}
	bird := Passaro{Nome: "Piu"}

	FazerAnimalAgir(dog)
	FazerAnimalAgir(bird)

	fmt.Println("\n=== Exemplo 6: Slice de interfaces ===")
	animais := []Animal{dog, bird}
	for _, animal := range animais {
		FazerAnimalAgir(animal)
	}
}
