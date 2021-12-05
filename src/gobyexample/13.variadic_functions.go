package gobyexample

import "fmt"

func Example13() {
	// Llamo de forma normal
	sum(1, 2, 3, 4, 5, 6)

	// Le paso un array o slice y uso el operador ...
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum(nums...)
}

// Toma N parametros del mismo tipo!
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, v := range nums {
		total += v
	}

	fmt.Println(total)
}

// Los argumentos variables se pueden combinar con normales, pero se pueden poner solo al final
func sum2(a, b, c string, nums ...int) {

}
