package gobyexample

import (
	"fmt"
	"math"
)

// Defino una interfaz con m'etodos
// Luego definos varios structs.... con sus metodos
// Esos metodos tienen q tener los mismos nombres que los declarados en la interfaz
// Para poder hacer uso de "polimorfismo" desde el main.

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Recibe una geometry y simplemente llama a sus funciones!
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Example19() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}