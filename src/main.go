package main

import (
	"fmt"
	pk "mypackage/src/mypackage" // alias es pk
)

func main() {
	var myCar pk.AutoPublic
	myCar.Marca = "Toyota"
	myCar.Anio = 2017

	fmt.Println(myCar)

	// prueba acceso a metodo privado
	// var myOtroCarro pk.autoPrivado
	// fmt.Println(myOtroCarro)

	pk.PrintMessage()

	// pk.printMessage()
}
