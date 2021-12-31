package gobyexample

import "fmt"

// Embedding: Se refiere a tener un Struct que wrappea a otro
// Se usa para simular Herencia

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprint("Numero ", b.num)
}

type container struct {
	base
	str string
}

func Example20() {

}