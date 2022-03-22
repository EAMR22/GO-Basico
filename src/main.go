package main

import "fmt"

func main() {
	// Declaracion de constantes
	// const pi float64 = 3.14
	// const pi2 = 3.1415
	// fmt.Println("pi:", pi) // fmt es una impresion de pantalla en la terminal.
	// fmt.Println("pi2:", pi2)

	// declaracion de variables enteras
	// base := 12 // El := Significa que esta variable no a sido declarada anteriormente.
	// var altura int = 14
	// var area int
	// fmt.Println(base, area, altura)

	// Zero values
	// var a int
	// var b float64
	// var c string
	// var d bool
	// fmt.Println(a, b, c, d)

	// Area cuadrado
	// const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println("El area del cuadrado es:", areaCuadrado)

	x := 50
	y := 10

	// Suma:
	result := x + y
	fmt.Println("Suma:", result)

	// Resta:
	result = x - y
	fmt.Println("Resta:", result)

	// Multiplicacion:
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Division:
	result = x / y
	fmt.Println("Division:", result)

	// Modulo:
	result = x % y
	fmt.Println("Modulo:", result)

	// Incrementa:
	x++
	fmt.Println("Incremental:", x)

	// Decremental:
	x--
	fmt.Println("Decremental:", x)
}
