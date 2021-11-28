package gobyexample

import "fmt"

func Example9() {
	// Son una vista a un array
	array := [4]int{9, 8, 7, 6} // Creo primero un array
	fmt.Println(array)

	slice := array[:2] // Le aplico un "slice operator" que crea un slice o vista a partir del elemento 0 al 2
	fmt.Println(slice, len(slice), cap(slice)) // cap es una Builtin function.. aunque el slice tenga len 2, el array por detras tiene mas elementos

	// Se pueden hacer slices directamente sin hacer antes un array:
	// Puedo crear un Slice de esta forma (de tipo int y con 5 de capacidad inicial):
	s := make([]string, 3) // Funcion Builtin
	// o podria ser:
	// s := make([]string, 3, 150) > Siendo 150 la capacidad
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// Get, set y len igual a los Arrays

	// Puede mutar de tamanio (en realidad crea uno nuevo y guardo lo que me retorna)
	s = append(s, "d") // Funcion Builtin
	s = append(s, "e", "f")
	fmt.Println(len(s))
	fmt.Println(s)

	// Puedo copiar un slice en otro (se copian sus valores)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Println(&s[5], &c[5]) // Con esto chequeo que realmente son dos punteros distintos, es, decir, el valor se copio

	// Operador: "slice" [inclusive : no inclusive]
	l := s[2:5]               // Crea un slice solo con los valores 2, 3, 4
	fmt.Println(&l[0], &s[2]) // Es el mismo puntero. El operador slice hace otro slice con los mismos elementos referenciados
	fmt.Println(l)

	l = s[:5] // Desde el inicio hasta el 5 no inclusive.
	fmt.Println(l)

	l = s[2:] // Desde el 2 hasta el final
	fmt.Println(l)

	// Declarar e inicializar un slice en una linea
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// Slice multidimensional
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
