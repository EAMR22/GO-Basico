package main

import "fmt"

func main() {

	// Defer:
	defer fmt.Println("Hola") // Con defer ejecutade ultimo esta linea de codigo.
	fmt.Println("Mundo")

	// Continue y Break:
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue:
		if i == 2 {
			fmt.Println("Continua el ciclo en", i)
			continue
		}

		// Break:
		if i == 7 {
			fmt.Println("Para el ciclo en", i)
			break
		}
	}
}
