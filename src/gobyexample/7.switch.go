package gobyexample

import "fmt"

func Example7() {
	// Como toda la vida
	// No es necesario el break, el default es opcional, se pueden tomar varios valores
	a := 4
	switch a {
	case 1:
		fmt.Println("Se puso 1")
	case 2:
		fmt.Println("Se puso 2")
	case 3, 4:
		fmt.Println("Se puso 3 o 4")
	default:
		fmt.Println("Se puso otra cosa")
	}

	// Como una alternativa al if (sin pasar valores)
	switch {
	case a > 6:
		fmt.Println("Es mayor a 6")
	default:
		fmt.Println("Es menor a 6")
	}

	// Para comparar tipos:
	whatAmI := func(i interface{}) string {
		switch t := i.(type) {
		case int:
			return "Procesando entero!"
		case bool:
			return "Procesando booleano"
		default:
			return fmt.Sprintf("No puedo procesar el tipo %T", t)
		}
	}
	fmt.Println(whatAmI(45))
	fmt.Println(whatAmI("Holaa"))
}
