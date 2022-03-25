package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// CarPrivate Car con acceso privado
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage es para imprimir un mensaje.
func PrintMessage(text string) {
	fmt.Println(text)
}
