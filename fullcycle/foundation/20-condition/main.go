package main

func SenioridadeIf(anos int) string {
	status := ""

	if anos < 3 {
		status = "Junior"
	} else if anos < 7 {
		status = "Pleno"
	} else {
		status = "Senior"
	}

	return status
}

func SenioridadeSwitch(anos int) string {
	status := ""

	switch {
	case anos < 3:
		status = "Junior"
	case anos < 7:
		status = "Pleno"
	default:
		status = "Senior"
	}

	return status
}

func main() {
	println(SenioridadeIf(9))
	println(SenioridadeSwitch(100))
}
