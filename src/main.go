package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	c <- "mensaje 1"
	c <- "mensaje 2"

	fmt.Println(len(c), cap(c)) // len() indica cuantas go routines hay en ese channel.
	// cap() nos indica cuanta es la cantidad maxima que puede almacenar en ese channel.

	// Range y Close:

	close(c) // Le indica al rountime de go que va a cerrar el canal.

	for message := range c {
		fmt.Println(message)
	}

	// Select:

	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1) // Imprime primero
	go message("mensaje2", email2) // Imprime despues

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1) // Imprime de ultimo
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2) // Imprime primero
		}
	}
}
