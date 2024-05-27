package main

import (
	"fmt"
)

func main() {
	// ciclo for
	for i := 0; i < 10; i++ {
		fmt.Println("valor de i:", i)
	}

	// ciclo for while
	contador := 0
	for contador < 10 {
		fmt.Println("Valor for while contador:", contador)
		contador++
	}

	// ciclo for forever
	// para terminar la ejecucion apreta control + c
	// for {
	// 	fmt.Println("Dalthon es un crack!")
	// }

	// for range, similar a foreach en otros lenguajes
	listaFrutas := []string{"pera", "manzana", "platano", "higo", "durazno"}

	for index, fruta := range listaFrutas {
		fmt.Printf("Fruta %d: %s\n", index, fruta)
	}
}
