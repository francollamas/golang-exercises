package gobyexample

import "fmt"

// Operador: & obtener el PUNTERO
// Operador: * obtener el VALOR

// No puedo cambiar valores de memoria a los punteros

func Example16() {
	i := 1
	fmt.Println("initial", i)

	// Se manda el valor por la funcion, pero llega como una copia de este, por eso no se actualiza
	changeval(i)
	fmt.Println("changeval", i)

	// Se manda el PUNTERO al valor, por lo que luego si se cambia.
	changeptr(&i)
	fmt.Println("changeptr", i)
}

func changeval(num int) {
	num = 5
}

func changeptr(num *int) {
	*num = 5 // Obtiene el "valor" del puntero, y se puede asignar.
}
