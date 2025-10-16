package inputs

import (
	"fmt"
	"strconv"
)

// GetMonth captura e valida o mês de nascimento
func GetMonth() string {
	var month string

	for {
		fmt.Println("------------------------------")
		fmt.Println("digite aqui o seu mês de nascimento... (ex: JAN = 1 | FEV = 2 etc...)")

		_, err := fmt.Scanln(&month)
		if err != nil {
			fmt.Println("Entrada inválida, por favor digite algum valor")
			continue
		}

		// Verifica o tamanho
		if len(month) < 1 || len(month) > 2 {
			fmt.Println("O mês deve ter 1 ou 2 digitos, digite novamente")
			continue
		}

		// Verifica se é um número
		monthInt, err := strconv.Atoi(month)
		if err != nil {
			fmt.Println("Por favor, digite apenas números para o mês")
			continue
		}

		// Verifica se está entre 1 e 12
		if monthInt < 1 || monthInt > 12 {
			fmt.Println("O mês deve ser entre 1 e 12, digite novamente")
			continue
		}

		break
	}

	return month
}
