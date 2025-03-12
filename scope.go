package main

import "fmt"

var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	var p int = 100
	var q = 200

	add(p, q)
	add(a, b)
	add(p, a)

	if a > 18 {
		l := 10000
		fmt.Println("This is a if block")
		fmt.Println("This print", l, " Value")
	}
}
