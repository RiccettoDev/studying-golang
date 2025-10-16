package inputs

import (
	"fmt"
	"strconv"
)

// GetYear captura e valida o ano de nascimento
func GetYear() string {
	var year string

	for {
		fmt.Println("------------------------------")
		fmt.Println("digite aqui o ano de seu nascimento... (ex:1999)")

		_, err := fmt.Scanln(&year)
		if err != nil {
			fmt.Println("Entrada inválida, por favor digite algum valor")
			continue
		}

		if len(year) != 4 {
			fmt.Println("O ano deve ter 4 digitos, digite novamente")
			continue
		}

		// Converte para número para validar
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			fmt.Println("Por favor, digite apenas números")
			continue
		}

		// Validação com números
		if yearInt < 1930 {
			fmt.Println("Mentira, este não é seu ano de nascimento, digite novamente")
			continue
		}

		if yearInt > 2024 {
			fmt.Println("Mentira, este não é seu ano de nascimento, digite novamente")
			continue
		}

		break
	}

	return year
}
