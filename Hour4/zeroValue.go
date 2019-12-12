package main

import "fmt"

var x int
var y string
var z float64
var a float32
var b = true

func main() {
	printvalues()
	types()
	Assignvalues()
	printvalues()
	types()
}
func types() {
	fmt.Println("\nLets see the types")
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("a is of type %T\n", a)
	fmt.Println("b is of type %T", b)
}
func Assignvalues() {
	fmt.Println("\n Assigning values")
	x = 21
	y = "I am a String"
	z = 2.123456789
	a = 1.12345678

}
func printvalues() {
	fmt.Println("\nPrinting the value of x =", x)

	fmt.Println("Printing the value of y =", y)

	fmt.Println("Printing the value of z =", z)

	fmt.Println("Printing the value of a =", a)
}
