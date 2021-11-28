package gobyexample

import "fmt"

func Example10() {

	// Creo con Make
	m := make(map[string]int, 50) // El size creo que esta de mas

	// Get y Set
	m["primero"] = 456
	m["segundo"] = 789
	fmt.Println(m["primero"])

	fmt.Println(m)

	// Builtin Detele function
	delete(m, "primero")

	fmt.Println(m)

	fmt.Println(len(m)) // Puedo calcular su tamanio

	// Indica si estaba presente el valor! esto distingue de keys que valen "" o 0
	_, prs := m["segundo"]
	fmt.Println(prs)

	// Crear e inicializar en la misma liena
	ma := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(ma)
}