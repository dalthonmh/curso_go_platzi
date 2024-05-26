package main

import "fmt"

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

}
