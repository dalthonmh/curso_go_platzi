package mypackage

import "fmt"

// AutoPublic Define la estructura inicial de un auto
type AutoPublic struct {
	Marca	string
	Anio 	int
}

type autoPrivado struct {
	marca string
	anio int
}

func PrintMessage() {
	fmt.Println("Hola desde mypackage")
}

func printMessage(){
	fmt.Println("Mensaje privado")
}