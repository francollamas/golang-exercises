package gobyexample

import "fmt"

func Example8() {

	// Creo un Array (no se puede cambiar su tamanio)
	var a [5]int
	fmt.Println("Arrays iniciados:", a)

	// Seteo y obtengo
	a[0] = 3
	fmt.Println(a[0])

	// Obtener el tamanio del array (funcion del lenguaje)
	fmt.Println("len:", len(a))

	// Declaro e inicio en una sola linea (la unica forma es con la Forma Corta de declarar variables)
	b := [3]int{1, 2, 3}
	fmt.Println(b)

	// Puedo hacer estructuras multidimensionales
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
