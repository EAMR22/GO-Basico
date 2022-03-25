package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() { // myPC actua como puntero
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a // Accedemos a la memoria de a.

	fmt.Println(b)
	fmt.Println(*b) // Accedemos al valor de la direccion de memoria.

	*b = 100
	fmt.Println(a) // Cambia el valor de "a" en 100.

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
