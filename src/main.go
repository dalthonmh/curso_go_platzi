package main

import "fmt"

func main() {
	// declaracion de variables
	helloMessage := "Hola"
	worldMessage := "mundo"

	// Println con fmt
	fmt.Println(helloMessage, worldMessage)

	persona := "Dalthon"
	actividad := "Curso de Go"
	anio := 2024
	// Donde: %s es para variables de tipo string y %d para variables de tipo int, %v es indistinto el tipo de dato
	// Le pasamos los valores en orden

	// Printf
	fmt.Printf("%s esta aprendiendo %s el anio %d\n", persona, actividad, anio)

	// Sprintf
	message := fmt.Sprintf("%s esta aprendiendo %s el anio %d", persona, actividad, anio)
	fmt.Println(message)

	// Para saber tipo de datos
	fmt.Printf("Tipo de dato de 'persona': %T\n", persona)
	fmt.Printf("Tipo de dato de 'anio': %T\n", anio)

}
