package main

import (
	"fmt"
)

func main() {
	fmt.Println("Switch Case")

	x := 11

	switch {
	case x < 10:
		fmt.Println("X é menor que 10")
	case x > 10:
		fmt.Println("X é maior que 10")
	case x == 10:
		fmt.Println("X é igual a 10")
	}

	// para cada case, pode ser colocado várias expressões para comparação
	// como por exemplo

	hojeQuemEstaNaFirma := "José"

	switch hojeQuemEstaNaFirma {
	case "José", "Maria":
		fmt.Println("Hoje quem está na firma é o time 1")
		fallthrough // fallthrough garante que o verificação condicional do próximo case não seja executada e execute direto a instrução
		// portanto com o fallthrough aqui o que seria exibido seria:
		// fmt.Println("Hoje quem está na firma é o time 1") e depois fmt.Println("Hoje quem está na firma é o time 2") também
	case "Carlos", "Amanda":
		fmt.Println("Hoje quem está na firma é o time 2")
	case "João", "Suzan":
		fmt.Println("Hoje quem está na firma é o time 3")
	default: // default para caso nenhuma opção se enquadrar
		fmt.Println("Hoje ninguem está na firma")
	}

	var y interface{}

	y = "batman"

	switch y.(type) {
	case string:
		fmt.Println("Y é do tipo string")
	case float64:
		fmt.Println("Y é do tipo float64")
	case bool:
		fmt.Println("Y é do tipo bool")
	case int:
		fmt.Println("Y é do tipo int")
	}

	// operadores condicionais:
	// ==   igual
	// !=   diferente
	// !    negação
	// %    resto da divisão
	// >    maior
	// <    menor
	// <=   menor igual
	// >=   maior igual
	// ||   OU
	// &&   E
	// =    atribuição
	// :=    atribuição curta
}
