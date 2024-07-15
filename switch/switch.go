package main

import "fmt"

func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func dayOfWeekTwo(day int) string {
	switch {
	case day == 1:
		return "Domingo"
	case day == 2:
		return "Segunda"
	default:
		return "Dia inválido"
	}
}

func dayOfWeekThree(day int) string {
	var dayOfWeek string
	switch {
	case day == 1:
		dayOfWeek = "Domingo"
		fallthrough // Manda pra proxima condição, sem avaliar.
	case day == 2:
		dayOfWeek = "Segunda"
	default:
		dayOfWeek = "Dia inválido"
	}

	return dayOfWeek
}

func main() {
	number := 1

	fmt.Println(dayOfWeek(number))
}
