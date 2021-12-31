package gobyexample

import "fmt"

func Example15() {
	fmt.Println(fact(7))

	// Fibonacci
	// Se puede hacer recursion con un closure, pero DEBE definirse su firma en una variable aparte para poder reutilizarse dentro.
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)

	}
	fmt.Println(fib(7))
}

// Factorial
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}