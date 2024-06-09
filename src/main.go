package main

import (
	"fmt"
)

func main() {
	peso_cohetes := make(map[string]int)

	peso_cohetes["starship"] = 5000000
	peso_cohetes["falcon9"] = 549054
	peso_cohetes["dragon"] = 4201
	peso_cohetes["falconheavy"] = 1420788

	fmt.Println(peso_cohetes)

	// recorrer valor
	for i, v := range peso_cohetes {
		fmt.Println(i, v)
	}

	// obtener un solo valor
	// ok indica si la llave existe o no en el diccionario
	value, ok := peso_cohetes["dragon"]
	fmt.Println(value, ok)

	// obtener un valor no existente
	value2, ok := peso_cohetes["noexiste"]
	fmt.Println(value2, ok)

	// otra forma de declarar
	//Declaration
	var map_1 = map[string]int32{
		"Car":      50000,
		"House":    20000,
		"Computer": 1000,
	}

	fmt.Println(map_1)
}
