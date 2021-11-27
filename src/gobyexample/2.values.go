package gobyexample

import "fmt"

func Example2() {
	// Concatenacion de string
	fmt.Println("Hola " + " Mundo")

	// Si divido enteros, el resultado es entero tambien
	fmt.Println(3 / 4)

	// Si divido decimales, el resultado es decimal
	fmt.Println(3 / 4.0)

	// Cualquier expresion primero se evalua, se resuelve, y luego se manda como parametro
	fmt.Println(true && false)

}
