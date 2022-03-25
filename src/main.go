package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // Acumula un conjunto de go routines y la va liberando poco a poco.

	fmt.Println("Hello")
	wg.Add(1) // Agregamos una go routine.

	go say("world", &wg)

	wg.Wait() // Le decimos al main que espere a las go routines de WaitGroup a que finalice.

	// Funcion anonima:

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1) // No es recomendable
}
