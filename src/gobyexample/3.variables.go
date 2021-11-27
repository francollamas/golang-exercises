package gobyexample

import "fmt"

func Example3() {
	// Todas las variables que no sean puntero, tienen valor por defecto y no admiten 'nil'

	// Declaro solamente.. esto asigna su valor por defecto: 0
	var a int
	fmt.Println(a)

	// Declaro y asigno al mismo tiempo (inferencia de tipos)
	var b = "Franco"
	fmt.Println(b)

	// Declaro (con tipo) y asigno al mismo tiempo
	var c string = "Emmanuel"
	fmt.Println(c)

	// Declaro varias (con o sin inferencia)
	var d, e string = "Pedro", "Miguel"
	fmt.Println(d, e)
	var f, g = "Jorge", "Luis"
	fmt.Println(f, g)

	// Forma corta de declarar e inicializar al mismo tiempo
	h := "Hola"
	fmt.Println(h)

	// Cualquiera de cualquier tipo se puede reasignar...
	a = 44
	h = "Chau"
	fmt.Println(a, h)
}
