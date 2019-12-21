package main

import "fmt"

var x int8 = 127
var a = -21
var b uint

func main() {
	y := 2
	z := 4.6548
	b = 255
	fmt.Println("\nThe Value of x is ", x)
	fmt.Println("The Value of y is ", y)
	fmt.Println("The Value of z is ", z)
	fmt.Println("The Value of a is ", a)
	fmt.Println("The Value of b is ", b)
	fmt.Println("\nLets see the types\n")
	fmt.Printf("x is of type %T \n", x)
	fmt.Printf("y is of type %T \n", y)
	fmt.Printf("z is of type %T \n", z)
	fmt.Printf("a is of type %T \n", a)
	fmt.Printf("b is of type %T \n", b)
}
