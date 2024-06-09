package main

import (
	"fmt"
)

func main() {
	// array y slices
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)
	fmt.Println(array, "tamanio array:", len(array), "capacidad array:", cap(array))

	// slices
	// en los slices no va la cantidad de tamaño que va a tener el array
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, "tamanio slice:", len(slice), "capacidad slice:", cap(slice))

	// slicing
	// metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])
}
