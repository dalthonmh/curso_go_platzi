package main

import (
	"fmt"
	"math"
)

func mandaSaludo(mensaje string) {
	fmt.Println(mensaje)
}

func tripleArgumento(a int, b int, c int) {
	fmt.Println(a, b, c)
}

func buenaPracticaTripleArgumento(a, b int, c string) {
	fmt.Println(a, b, c)
}

func retornaValor(a int) int {
	return a * 2
}

func retornaDosValores(a int) (c, d int) {
	return a, a * 3
}

// area del circulo
func calculaAreaCirculo(radio float64) float64 {
	return math.Pi * math.Pow(radio, 2)

}

// area del trapecio
func calculaAreaTrapecio(baseMayor, baseMenor, altura float64) float64 {
	return (baseMayor + baseMenor) * altura / 2
}

func main() {
	mandaSaludo("Hola mundo")
	tripleArgumento(1, 2, 3)
	buenaPracticaTripleArgumento(1, 2, "lalo")

	valor := retornaValor(2)
	fmt.Println("El valor retornado es:", valor)

	valor1, valor2 := retornaDosValores(3)
	fmt.Println("Valor1 y valor2 son:", valor1, valor2)

	dato1, _ := retornaDosValores(2)
	fmt.Println("Solo uso 1er valor:", dato1)

	_, dato2 := retornaDosValores(5)
	fmt.Println("Solo uso 2do valor:", dato2)

	circulo := 3.4
	areaCirculo := calculaAreaCirculo(circulo)
	fmt.Println("Area circulo: ", areaCirculo)

	base_mayor := 6.3
	base_menor := 3.2
	altura := 2

	areaTrapecio := calculaAreaTrapecio(base_mayor, base_menor, float64(altura))
	fmt.Println("Area del trapecio es: ", areaTrapecio)
}
