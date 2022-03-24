package main

import "fmt"

func main() {

	// Array:
	var array [4]int
	array[0] = 1
	array[1] = 2

	// Con "len" le dice cuantos elementos hay en el array.
	// Con "cap" nos dice la capacidad maxima del array.
	fmt.Println(array, len(array), cap(array))

	// Slice:
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos del slice:
	fmt.Println(slice[0])   // Imprime el elemento 0.
	fmt.Println(slice[:3])  // Imprime desde el inicio hasta el 2, el 3 no porque es exclusivo.
	fmt.Println(slice[2:4]) // Imprime del 2 hasta el 3, el 4 no porque es exclusivo.
	fmt.Println(slice[4:])  // Imprime del 4 en adelante, va desde el 4 porque es inclusivo.

	// Append:
	slice = append(slice, 7) // Con append agregaun elemento a la lista.
	fmt.Println(slice)

	// Append nueva lista:
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Aqui une las 2 listas.
	fmt.Println(slice)
}
