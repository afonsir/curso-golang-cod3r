package main

import "fmt"

func main() {
	employeesByletter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Júnior": 1200.00,
		},
	}

	delete(employeesByletter, "P")

	for letter, employees := range employeesByletter {
		fmt.Println(letter, employees)
	}
}
