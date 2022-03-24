package main

import "fmt"

func main() {

	m := make(map[string]int) // Con map estamos guardando un diccionario.

	m["Jose"] = 14
	m["Maria"] = 20

	fmt.Println(m)

	// Recorrer map:
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor:
	value, ok := m["Maria"] // "ok" nos indica si un valor existe en el diccionario.
	fmt.Println(value, ok)
}
