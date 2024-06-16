package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

// methods
func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// recibe la interfaz y lo imprime
func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	miCuadrado := cuadrado{base: 5}
	miRectangulo := rectangulo{base: 3, altura: 5}

	calcular(miCuadrado)
	calcular(miRectangulo)

	// lista de interfaces
	miIntefaz := []interface{}{"Diferente", 4, true, 3.54, "de datos"}
	fmt.Println(miIntefaz...)
}
