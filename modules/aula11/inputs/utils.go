package inputs

import "strconv"

// IsLeapYear verifica se um ano é bissexto
func IsLeapYear(yearStr string) bool {
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return false
	}

	// Regra do ano bissexto:
	// - Divisível por 4
	// - Mas não por 100, a menos que também seja divisível por 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// GetDaysInMonth retorna o número de dias em um mês específico
func GetDaysInMonth(monthStr, yearStr string) int {
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if IsLeapYear(yearStr) {
			return 29
		}
		return 28
	default:
		return 0
	}
}
