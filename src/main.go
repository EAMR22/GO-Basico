package main

import "fmt"

// Structs: La forma de hacer clases con go.

type car struct { // Con "type" nos indica que es un tipo de dato.
	brand string
	year  int
}

func main() {

	// Instanciamos un opjeto:
	myCar := car{brand: "Porche", year: 2008}
	fmt.Println(myCar)

	// Otra forma:
	var otherCar car
	otherCar.brand = "Lamborghini"
	fmt.Println(otherCar)
}
