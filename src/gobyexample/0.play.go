package gobyexample

import "fmt"

type Person struct {
	name string
	age int
}

type Franco string

// Hago override del metodo String
func (p Person) String() string {
	return "asd"
}

func Play() {
	person := Person{
		"Franco", 45,
	}

	fmt.Println(person)

	var asd Franco
	asd = "asdasd"
	fmt.Println(asd)
	dsa := fmt.Sprintf("Hola %b", asd)
	fmt.Println(dsa)

	aaa := 45

	fmt.Printf("Hola %s", aaa)

}
