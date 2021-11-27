package gobyexample

import "fmt"

func Example5() {
	// El analogo a un 'while'
	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}

	// El clasico
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	// Infinito (hasta que se haga break o return)
	for {
		fmt.Println("\nSe ejecuta")
		break
	}

	// Clasico, con uso de sentencia "continue"
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}

	// Luego hay mas tipos de for cuando veamos rangos, canales y otras estructuras.

}
