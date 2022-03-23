package main

import "fmt"

func normalFuntion(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) { // Asi es una buena practica en go.
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // Retorna un valor
	return a * 2
}

func doubleReturn(a int) (c, d int) { // retorna 2 valores
	return a, a * 2
}

func main() {
	normalFuntion("Hola mundo")
	tripeArgument(1, 4, "vamossss")

	value := returnValue(2)
	fmt.Println("value:", value)

	value1, _ := doubleReturn(2) // El "_" hace que no tome en cuenta el segundo valor.
	fmt.Println("value1:", value1)
}
