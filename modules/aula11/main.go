package main

import (
	"fmt"
	"go-api/modules/aula11/inputs"
)

func main() {
	fmt.Println("=== CALCULADORA DE IDADE ===")

	// Primeiro pega o ano e mês para validar o dia depois
	year := inputs.GetYear()
	month := inputs.GetMonth()
	day := inputs.GetDay(month, year)

	fmt.Println("\n=== RESUMO ===")
	fmt.Printf("Data de nascimento: %s/%s/%s\n", day, month, year)

	if inputs.IsLeapYear(year) {
		fmt.Println("⭐ Este é um ano bissexto!")
	}
}
