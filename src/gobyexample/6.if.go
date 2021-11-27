package gobyexample

import "fmt"

func Example6() {
	// Los If y else funcionan como en cualquier otro lenguaje

	// Se pueden declarar variables en el mismo if, y son accesibles en cualquier bloque dentro del if
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// No existe el operador ternario

}
