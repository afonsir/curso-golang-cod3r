package main

import "fmt"

func getResult(grade float64) string {
	if grade >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(getResult(6.5))
}
