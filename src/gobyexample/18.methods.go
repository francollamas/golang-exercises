package gobyexample

import "fmt"

type rect struct {
	width, height int
}

// Pueden ser definidos por puntero o valor
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) cambiar() {
	r.width = 1
}

func (r *rect) cambiarConPtr() {
	r.width = 1
}


func Example18() {
	r := rect{34, 45}
	fmt.Println(r)

	r.cambiar() // No cambia nada xq se llamo x referencia y como q se hizo una copia...
	fmt.Println(r)

	r.cambiarConPtr() // Si cambia, ya q se le paso un puntero!
	fmt.Println(r)
}