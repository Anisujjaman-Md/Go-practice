package main

import (
	"fmt"

	"example.com/mathlib"
)

var (
	a = 20
	b = 40
)

func addTwoNumber(x int, y int) {
	res := x + y
	printRes(res)
}

func main() {
	mathlib.Add(a, b)
	addTwoNumber(a, b)

	//IIFE anonymous func

	func(a, b int) {
		k := a + b
		fmt.Println(k)
	}(1, 2)

	//Noob Function

	mul := func(c, d int) {
		l := c * d
		fmt.Println(l)
	}

	mul(10, 2)

}

func printRes(res int) {
	fmt.Println("This is Final Result", res)
}

func init() {
	fmt.Println("This is init function")
}
