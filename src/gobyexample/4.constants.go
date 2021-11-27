package gobyexample

import (
	"fmt"
	"math"
)

const s string = "3"

const (
	x, y = 6, 7
	z = 55
)

const (
	w = 33
)

func Example4() {

	fmt.Println(w, x, y, z)

	// Pueden ser definidas en los mismos lugares que las variables.
	const b = "Franco"

	// Una constante numerica no tiene tipo hasta que se le da uno, por ejemplo a traves de una conversion
	const c = 90
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", int64(c))

	// Notacion cientifica (son float64)
	const d = 2e-20
	const e = 2e45
	fmt.Println(d)
	fmt.Printf("%T\n", e)

	// Una CONSTANTE numerica SIN tipo (no variables) puede tomar un tipo especifico segun lo que requiera su contexto
	// Ej: la funcion Sin pide un float64, pero aqui le estamos pasando una constante numerica sin tipo.. esta se autoconvierte a float64
	math.Sin(c)
	fmt.Println(c)
}
