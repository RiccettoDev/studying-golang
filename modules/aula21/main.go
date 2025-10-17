package main

import (
	"fmt"
	funccalc "go-api/modules/aula21/func_calc"
)

func main() {
	fmt.Println("==========================")
	fmt.Println(".   Calculadora GoLang   .")
	fmt.Println("==========================")

	for {
		fmt.Println("")
		fmt.Println("O que gostaria de fazer?")
		fmt.Println("")
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtrair")
		fmt.Println("3 - Multiplicar")
		fmt.Println("4 - Dividir")
		fmt.Println("0 - Sair")

		var option int
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&option)

		if option == 0 {
			fmt.Println("\nEncerrando a calculadora... 👋")
			break
		}

		if option < 0 || option > 4 {
			fmt.Println("\nOpção inválida! Tente novamente.")
			continue
		}

		var num1, num2 float64
		fmt.Println("\nDigite o primeiro número:")
		fmt.Scanln(&num1)
		fmt.Println("Digite o segundo número:")
		fmt.Scanln(&num2)

		switch option {
		case 1:
			result := funccalc.Sum(num1, num2)
			fmt.Printf("\nO resultado da soma é: %.2f\n\n", result)

		case 2:
			result := funccalc.Subtraction(num1, num2)
			fmt.Printf("\nO resultado da subtração é: %.2f\n\n", result)

		case 3:
			result := funccalc.Multiplication(num1, num2)
			fmt.Printf("\nO resultado da multiplicação é: %.2f\n\n", result)

		case 4:
			if num2 == 0 {
				fmt.Println("\nErro: divisão por zero não é permitida.")
				continue // volta para o início do loop
			}
			result := funccalc.Division(num1, num2)
			fmt.Printf("\nO resultado da divisão é: %.2f\n\n", result)

		default:
			fmt.Println("Opção inválida! Tente novamente.")
		}
	}
}
