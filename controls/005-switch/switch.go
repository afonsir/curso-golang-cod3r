package main

import "fmt"

func gradeToConcept(n float64) string {
	var grade = int(n)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(gradeToConcept(9.8))
	fmt.Println(gradeToConcept(6.9))
	fmt.Println(gradeToConcept(2.1))
}
