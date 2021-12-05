package gobyexample

import "fmt"

// Funciones anonimas: para definir funciones sin tener que nombrarlas
// Closures: funcion dentro de otra... que puede usar una variable de una funcion externa.

func Example14() {
	// Guardo la funcion en una variable
	nextInt := closureFunction()

	// La llamo varias veces y el valor de la 'i' de la funcion externa va variando!
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt2 := closureFunction()
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())
	fmt.Println(nextInt2())

	// Solo jugando con funciones anonimas...
	experimentando()
}

// Seria como una funcion con estado, ya que al retornarla sigue manteniendo el estado de todas las variables
// declaradas en la que esta por fuera
func closureFunction() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Experimento creando una funcion anonima que recibe parametros, y la paso a otra funcion para que esta la ejecute.
// Los estados se siguen manteniendo por mas que esta funcion se pase por todos lados.
func experimentando() {

	name := "Franco"
	funcion := func(a string, b int) string {
		name += "o"
		return fmt.Sprintln("Hola", name, a, b)
	}

	funcion2(funcion)
}

// Paso funcion por parametro
func funcion2(x func(string, int) string) {
	fmt.Println(x("aa", 4))
	fmt.Println(x("bb", 5))
	fmt.Println(x("cc", 6))
	fmt.Println(x("dd", 7))
}