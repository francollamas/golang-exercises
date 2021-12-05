package gobyexample

import "fmt"

func func1() {
	fmt.Println("Test func1")
}

func func2(a int, b int) int {
	return a + b
}

// Omito el tipo para parametros consecutivos
func func3(a, b, c int, d, e string) int {
	fmt.Println(d, e)
	return a + b + c
}

// Funcion con retorno multipple
func func4(a, b int) (int, int) {
	return a + b, a - b
}

func Example12() {
	func1()

	a := func2(2, 3)
	fmt.Println(a)

	b := func3(2, 3, 4, "hola", "mundo")
	fmt.Println(b)

	c, d := func4(7, 5)
	fmt.Println(c, d)

	// Dismiseo la que no me interesa
	e, _ := func4(7, 5)
	fmt.Println(e)
}