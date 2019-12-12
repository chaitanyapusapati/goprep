package main

import "fmt"

var x, z int
var y = "string"

func main() {
	w := 12 + 10
	fmt.Println(w)
	x = 42
	y = "yes, It is a string"
	foo()

}
func foo() {
	fmt.Println("I am in foo")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println(z)

}
