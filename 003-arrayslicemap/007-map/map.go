package main

import "fmt"

func main() {
	approved := make(map[int]string)

	approved[12345678978] = "Maria"
	approved[23473987324] = "Pedro"
	approved[23409823809] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[23409823809])
	delete(approved, 23409823809)
	fmt.Println(approved[23409823809])
}
