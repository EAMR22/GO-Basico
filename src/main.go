package main

import "fmt"

func main() {

	// For condicional:
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For while:
	counter := 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}

	fmt.Printf("\n")

	// For forever:
	counterForever := 0 // Es un ciclo infinito
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
