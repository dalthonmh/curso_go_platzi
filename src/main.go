package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "mensaje1"
	c <- "mensaje2"

	fmt.Println(len(c), cap(c))

	// range y close
	close(c)
	// lo ideal es cerrar el canal cuando ya no lo utlicemos
	// c <- "mensaje3"

	for message := range c {
		fmt.Println(message)
	}

	// select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de 2", m2)

		}
	}
}
