package gobyexample

import "fmt"

type person struct {
	name string
	age  int
}

// En los constructures devuelvo PUNTEROS
func newPerson(name string) *person {
	p := person{
		name: name,
	}
	p.age = 42

	return &p
}

// Al usar el . (punto) para acceder a un campo de un tipo Puntero, es automaticamente derreferenciado.
// La diferencia sera que si paso x parametro un VALOR por funcion, del otro lado se hace una DEEP COPY. En cambio si paso un puntero no.

func Example17() {
	fmt.Println(person{"Franco", 25})
	fmt.Println(person{name: "Franco", age: 25})
	fmt.Println(person{age: 25})
	fmt.Println(person{name: "Franco"})
	fmt.Println(person{})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	sp.age = 45
	fmt.Println(sp)

	s.age = 46
	fmt.Println(s)

	recibirPersona(s)
	fmt.Println(s)
}

func recibirPersona(persona person) {
	persona.age = 11
	fmt.Println(persona)
}