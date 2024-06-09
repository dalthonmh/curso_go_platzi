package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string) {

	var textReverse string

	for i := len(texto) - 1; i >= 0; i-- {
		textReverse += string(texto[i])
	}

	if strings.ToLower(texto) == strings.ToLower(textReverse) {
		fmt.Println(texto, "==> SI es palindromo (OK)")
	} else {
		fmt.Println(texto, "==> No es palindromo")
	}

}

func main() {
	slice := []string{"lunes", "martes", "miercoles", "jueves", "viernes"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// recorrido sin indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// ----- detectar si es palindromo -----
	isPalindromo("Ama")
}
