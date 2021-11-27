package gobyexample

import (
	"fmt"
	"os"
)

func Example1() {
	// ******* PRINT *******
	// Imprime en la salida estandar
	// Recibe un array de tipo interface{}, es decir, le puedo mandar cualquier cosa
	// Concatena los parametros con espacio, y agrega un \n al final
	fmt.Println("Hello world!", "im", "Franco", "and im", 25, "years old")


	// ******* SPRINT *******
	// Igual pero en vez de imprimir en consola, lo devuelve como un string
	text := fmt.Sprintln("Hi!")
	text += fmt.Sprintln("im Franco")
	fmt.Println(text)


	// ******* FPRINT *******
	// Imprime en la salida indicada por parametro
	// os.Stdout es un File, estimo que de esta forma se podria escribir en otros archivos.
	fmt.Fprintln(os.Stdout, "Hello World!")


	// ******* Variaciones de PRINT *******
	// Igual que Println pero sin salto de linea \n
	fmt.Print("Print test")
	fmt.Print("Another print test\n")

	// Igual que Println pero sin salto de linea \n y permite setear formato
	// Devuelve la cantidad de bytes escritos y un error, generalmente no se usan
	age := 25
	bytes, err := fmt.Printf("Hi, im %s and im %v \n", "Franco", age)
	fmt.Println(bytes, err)

	// ******* ERRORF *******
	// Devuelve un error, no manda nada a ninguna salida.
	name := "Franco"
	err = fmt.Errorf("User %s not found", name)
	fmt.Println(err.Error())


	// ****** SCAN *******
	// Escanea desde la consola estandard.... se le deben pasar punteros a las variables
	// Lo que venga separado como espacio va a una nueva variable, no importa si hay saltos de linea, se ignoran.
	var scan1, scan2, scan3, scan4 string
	fmt.Scan(&scan1)
	fmt.Scan(&scan2)
	fmt.Scan(&scan3, &scan4)

	fmt.Println("Scan:", scan1, scan2)
	fmt.Println("Scan:", scan3, scan4)

	// Hace un Scan pero sabe exactamente c'omo obtener el valor. El texto que entra por la entrada estandar
	// debe ser igual al que esta especificado en el formato
	fmt.Scanf("Hola %s tenes %s anios" , &scan1, &scan2)
	fmt.Println("Scanf:", scan1, scan2)

	// Escanea hasta que se ingrese un enter!
	fmt.Scanln(&scan1, &scan2, &scan3)
	fmt.Println("Scanln", scan1, scan2, scan3)

	// Escanea desde un string y no desde la entrada estandard
	fmt.Sscan("Hola me llamo Franco", &scan1, &scan2, &scan3, &scan4)
	fmt.Println("Sscan:", scan1, scan2, scan3, scan4)

	// Igual que el anterior, pero con formato
	fmt.Sscanf("Hola me llamo Emmanuel", "Hola me llamo %s", &scan1)
	fmt.Println("Sscanf", scan1)

	// Igual que scan desde string, pero hasta que haya una nueva linea
	// En este caso aunque requiramos 3 valores, solo lee los dos primeros antes del \n y corta.
	fmt.Sscanln("Line one\nLine two", &scan1, &scan2, &scan3)
	fmt.Println("Sscanln:", scan1, scan2, scan3)

	// Escanea desde el File que le especifiquemos.
	fmt.Fscan(os.Stdin, &scan1)
	fmt.Println("Fscan:", scan1)
}