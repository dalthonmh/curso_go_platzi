package main

import "fmt"

type computadora struct {
	ram   int
	disk  int
	brand string
}

func (miComputadora *computadora) duplicarRam() {
	miComputadora.ram = miComputadora.ram * 2
}

func (miComputadora computadora) String() string {
	return fmt.Sprintf("Tengo %d GB de ram, %d GB de disco y es de marca %s", miComputadora.ram, miComputadora.disk, miComputadora.brand)
}

func main() {
	lenovo := computadora{ram: 4, disk: 200, brand: "Lenovo"}

	fmt.Println(lenovo)

	lenovo.duplicarRam()
	fmt.Println(lenovo)

	lenovo.duplicarRam()
	fmt.Println(lenovo)

}
