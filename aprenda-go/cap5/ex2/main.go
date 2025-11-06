// Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.

// ==  igual
// !=  diferente
// <=  menor igual
// >=  maior igual
// <   menor
// >   maior

// Demonstre estes valores.

package main

import (
	"fmt"
)

func main() {

	igual1 := 10 == 10
	igual2 := 5 == 10
	diferente1 := 10 != 10
	diferente2 := 5 != 10
	menor_igual1 := 5 <= 5
	menor_igual2 := 6 <= 5
	menor_igual3 := 2 <= 5
	maior_igual1 := 2 >= 2
	maior_igual2 := 6 >= 2
	maior_igual3 := 2 >= 5
	menor1 := 2 < 2
	menor2 := 1 < 2
	menor3 := 3 < 2
	maior1 := 2 > 2
	maior2 := 1 > 2
	maior3 := 3 > 2

	fmt.Printf("10 é igual a 10: %v\n", igual1)
	fmt.Printf("5 é igual a 10: %v\n", igual2)

	fmt.Println("==================")

	fmt.Printf("10 é diferente a 10: %v\n", diferente1)
	fmt.Printf("5 é diferente a 10: %v\n", diferente2)

	fmt.Println("==================")

	fmt.Printf("5 é menor igual a 5: %v\n", menor_igual1)
	fmt.Printf("6 é menor igual a 5: %v\n", menor_igual2)
	fmt.Printf("2 é menor igual a 5: %v\n", menor_igual3)

	fmt.Println("==================")

	fmt.Printf("2 é maior igual a 2: %v\n", maior_igual1)
	fmt.Printf("6 é maior igual a 2: %v\n", maior_igual2)
	fmt.Printf("2 é maior igual a 5: %v\n", maior_igual3)

	fmt.Println("==================")

	fmt.Printf("2 é menor que 2: %v\n", menor1)
	fmt.Printf("1 é menor que 2: %v\n", menor2)
	fmt.Printf("3 é menor que 2: %v\n", menor3)

	fmt.Println("==================")

	fmt.Printf("2 é maior que 2: %v\n", maior1)
	fmt.Printf("1 é maior que 2: %v\n", maior2)
	fmt.Printf("3 é maior que 2: %v\n", maior3)
}
