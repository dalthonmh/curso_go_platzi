package main

import (
	"fmt"
)

type auto struct {
	marca string
	anio  int
}

func main() {
	miauto := auto{marca: "toyota", anio: 2017}
	fmt.Println(miauto)

	// otra declaracion de auto
	var segundo_auto auto
	segundo_auto.marca = "tesla"

	fmt.Println(segundo_auto)
}
