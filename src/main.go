package main

import (
	"fmt"
)

func main() {
	// defer: sirve para ejecutar al ultimo de las instrucciones
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		//continue
		if i == 2 {
			fmt.Println("OJO: llegue a 2, continuo y no escribo el número de iteación")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("OJO: llegue a 8, rompo todo el bucle")
			break
		}
		fmt.Println("iteración", i, "terminada")
	}
}
