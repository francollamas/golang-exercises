package gobyexample

import "fmt"

func Example11() {

	// Range de Arrays y Slices (obtengo Index y Value)
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	// Range de Maps (obtengo Key y Value)
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	// Itero solo la Key del Map (obtengo Key)
	for k := range kvs {
		fmt.Println(k)
	}

	// Range en Strings (obtengo Index y Character > en numero)
	for i, c := range "Franco" {
		fmt.Println("index", i, "character", c)
	}
}