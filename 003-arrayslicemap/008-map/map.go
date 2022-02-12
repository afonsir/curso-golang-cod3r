package main

import "fmt"

func main() {
	employees := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	employees["Rafael Filho"] = 1350.0
	delete(employees, "Inexistente")

	for name, salary := range employees {
		fmt.Println(name, salary)
	}
}
