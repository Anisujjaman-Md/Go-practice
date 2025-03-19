package main

import (
	"fmt"
)

func add(x, y int) int {
	z := x + y
	fmt.Println(z)
	return z
}

func div(p, q int) int {
	r := p / q
	fmt.Println(r)
	return r
}

func PrintSomething(result int) {
	fmt.Println("The Output:", result)
}

// Higher-order function: accepts a function
func ExampleOfFuntionAsParemeter(a, b int, result func(e, f int) int) {
	result(a, b)
}

// Higher-order function: returns a function
func ExampleOfFuntionReturnOfAFunction() func(x, y int) int {
	return add
}

// Higher-order function: accepts a function and returns another function
func ExampleOfBoth(a, b int, operation func(int, int) int) func(int) {
	return func(result int) {
		PrintSomething(operation(a, b))
	}
}

func main() {
	ExampleOfFuntionAsParemeter(7, 8, add)
	ExampleOfFuntionAsParemeter(70, 10, div)

	res := ExampleOfFuntionReturnOfAFunction()

	res(1, 2)

	printFunc := ExampleOfBoth(3, 5, add)

	// Call the returned function with any int (not used)
	printFunc(0)
}
