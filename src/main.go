package main

import (
	"fmt"
)

func determinarParImpar(numero int) string {
	if numero%2 == 0 {
		return "Par"
	}
	return "Impar"
}

// Definimos una estructura de usuario
type Usuario struct {
	usuario  string
	password string
}

// Definimos la funcion para saber si existe el usuario
func login(usuario, password string) string {
	// Creamos un array de 3 elementos de usuarios
	var usuarios [3]Usuario

	// Asignamos valores a los elementos del array
	usuarios[0] = Usuario{usuario: "Dalthon", password: "123"}
	usuarios[1] = Usuario{usuario: "Pepe", password: "abc"}
	usuarios[2] = Usuario{usuario: "Jose", password: "xyz"}

	// Retorna token de acceso
	token := ""

	for _, user := range usuarios {
		if user.usuario == usuario && user.password == password {
			token = "token-" + user.usuario
		}
	}

	if token != "" {
		return token
	}

	return "No existe el usuario"

}

func main() {
	// Reto1: Determinar si un número es par o impar
	numero := 3
	respuesta := determinarParImpar(numero)
	fmt.Println("El número", numero, "es:", respuesta)

	// Reto2: Funcionalidad login
	usuario := "Dalthon"
	password := "123"

	isLoginOk := login(usuario, password)
	fmt.Println(isLoginOk)

}
