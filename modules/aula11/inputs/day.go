package inputs

import (
	"fmt"
	"strconv"
)

// GetDay captura e valida o dia de nascimento considerando o mês e ano
func GetDay(monthStr, yearStr string) string {
	var day string

	for {
		fmt.Println("------------------------------")
		fmt.Println("digite aqui o seu dia de nascimento... (ex:1, 2, 3, etc...)")

		_, err := fmt.Scanln(&day)
		if err != nil {
			fmt.Println("Entrada inválida, por favor digite algum valor")
			continue
		}

		if len(day) < 1 || len(day) > 2 {
			fmt.Println("O dia deve ter 1 ou 2 digitos, digite novamente")
			continue
		}

		// Converte para número para validar
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			fmt.Println("Por favor, digite apenas números")
			continue
		}

		// Validação básica
		if dayInt < 1 || dayInt > 31 {
			fmt.Println("Dia inválido! Digite um dia entre 1 e 31")
			continue
		}

		// Validação avançada: verifica se o dia existe no mês/ano
		maxDays := GetDaysInMonth(monthStr, yearStr)
		if dayInt > maxDays {
			monthInt, _ := strconv.Atoi(monthStr)
			monthNames := []string{
				"", "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
				"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
			}

			if monthInt == 2 && IsLeapYear(yearStr) {
				fmt.Printf("Fevereiro de %s tem 29 dias (ano bissexto)! Digite novamente\n", yearStr)
			} else {
				fmt.Printf("%s tem %d dias! Digite um dia entre 1 e %d\n",
					monthNames[monthInt], maxDays, maxDays)
			}
			continue
		}

		break
	}

	return day
}
