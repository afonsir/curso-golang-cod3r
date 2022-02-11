package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// int
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// unsigned int, uint8 uint16 uint32 uint64
	var b byte = 255 // uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// signed int
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // int32 (unicode)
	fmt.Println("O rune é do tipo", reflect.TypeOf(i2))
	fmt.Println(i2)

	// decimal
	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Afonso"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// multi-line string
	s2 := `Olá
	meu
	nome
	é
	Afonso`
	fmt.Println("O tamanho da string é", len(s2))

	// var x char = 'b'
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
