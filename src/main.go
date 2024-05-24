package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.15

	fmt.Println("pi :", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base: ", base)
	fmt.Println("altura: ", altura)
	fmt.Println("area: ", area)

	// valores por defecto si no los inicializamos (zero values)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)
}
