package main

import "fmt"

var x int
var y = "String"
var z float32
var a float64
var b bool

type own int

var c own

func main() {
	x = 535873
	z = 1.123456789012345678901234567890
	a = 1.123456789012345678901234567890
	c = 21
	values()
	conversion()
}
func values() {
	fmt.Println("\n   Lets See The Values\n")
	fmt.Println("Value of x is ", x)
	fmt.Println("Value of y is ", y)
	fmt.Println("Value of z is ", z)
	fmt.Println("Value of a is ", a)
	fmt.Println("Value of b is ", b)
	fmt.Println("Value of c is ", c)
	types()
}
func types() {
	fmt.Println("\n   Lets See  The Types\n")
	fmt.Printf("x is of Type %T\n", x)
	fmt.Printf("y is of Type %T\n", y)
	fmt.Printf("z is of Type %T\n", z)
	fmt.Printf("a is of Type %T\n", a)
	fmt.Printf("b is of Type %T\n", b)
	fmt.Printf("c is of Type %T\n", c)

}
func conversion() {
	fmt.Println("\n  Doing Conversion Not Casting\n")
	z = float32(a)
	a = float64(z)
	x = int(c)
	values()

}
