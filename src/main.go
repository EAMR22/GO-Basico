package main

import "fmt"

func main() {
	// Declaracion de variables:
	helloMessage := "hello"
	worldMessage := "world"

	// Println:
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf:
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)
	// %s recibe un string, %d recibe un entero, %v cuando no se sabe que tipo de dato es, \n es un salto de linea.

	//Sprintf:
	message := fmt.Sprintf("%s tiene mas de %d cursos\n", nombre, cursos) // Con Sprintf guarda la cadena en message
	fmt.Println(message)

	//Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage) // Con %T imprime el tipo de dato de la impresion.
	fmt.Printf("Cursos: %T\n", cursos)
}
