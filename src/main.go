package main

import (
	"fmt"
)

func main() {

	// con condicion
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Es par")

	default:
		fmt.Println("Es impar")
	}

	// sin condicion
	value := 50
	switch {
	case value > 50:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Printf("%d Es menor a 100 y mayor a 0\n", value)
	}

}
