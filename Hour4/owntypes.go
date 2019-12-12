package main

import "fmt"

var y int

type own int

var x own

func main() {
	y = 42
	fmt.Println(y)
	fmt.Printf("y is of type %T\n", y)
	x = 44
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)
	foo()
}
func foo() {
	type types int
	var z types
	z = 47
	fmt.Println(z)
	fmt.Printf("z is of type %T\n", z)
}
