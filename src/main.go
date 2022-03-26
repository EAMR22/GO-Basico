package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text // Le indicamos que vamos a ingresar un dato a ese canal.
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bay", c)

	fmt.Println(<-c) // La salida del dato del canal.
}
