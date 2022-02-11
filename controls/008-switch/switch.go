package main

import (
	"fmt"
	"time"
)

func kind(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(kind(2.3))
	fmt.Println(kind(1))
	fmt.Println(kind("Opa"))
	fmt.Println(kind(func() {}))
	fmt.Println(kind(time.Now()))
}
