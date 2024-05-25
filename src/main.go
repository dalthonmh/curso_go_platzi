package main

import (
	"fmt"
	"math"
)

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
	fmt.Println("-----------------")

	// valores por defecto si no los inicializamos (zero values)
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	// Area cuadrado
	// const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println(areaCuadrado)

	// operaciones aritmeticas
	x := 10
	y := 20
	// valores iniciales
	fmt.Println("Valor inicial de x: ", x)
	fmt.Println("Valor inicial de y: ", y)

	result := x + y

	// suma
	fmt.Println("suma: ", result)

	// resta
	result = y - x
	fmt.Println("resta: ", result)

	// multiplicacion
	result = x * y
	fmt.Println("producto: ", result)

	// division
	result = y / x
	fmt.Println("division: ", result)

	// incrementar
	x++
	fmt.Println("Incrementar x: ", x)

	// decrementar
	x--
	fmt.Println("decrementar x: ", x)

	fmt.Println("-----------------")
	// cacular areas
	// trapecio y circulo
	trapecio_base_mayor := 10
	trapecio_altura := 4
	trapecio_base_menor := 8

	fmt.Println("trapecio_base_mayor: ", trapecio_base_mayor)
	fmt.Println("trapecio_base_menor: ", trapecio_base_menor)
	fmt.Println("trapecio_altura: ", trapecio_altura)

	// formula de area es A = (a+b)h/2
	trapecio_area := (trapecio_base_mayor + trapecio_base_menor) * trapecio_altura / 2
	fmt.Println("Area de trapecio: ", trapecio_area)

	// area circulo
	fmt.Println("-----------------")
	circulo_radio := 4
	// formula area de circulo: A = pi * r * r
	fmt.Println("Radio de circulo: ", circulo_radio)

	circulo_area := pi * float64(circulo_radio) * float64(circulo_radio)
	fmt.Println("Area de circulo: ", circulo_area)
	fmt.Println("El área del círculo es:", math.Pi*math.Pow(float64(circulo_radio), 2))

}
