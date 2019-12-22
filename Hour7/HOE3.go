/* Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants. */
package main

import "fmt"

const (
	x int8    = 24
	y float32 = 2.74674
	z string  = "James Bond"
)
const (
	a = 44
	b = 4.65816
	c = "Bond"
)

func main() {
	fmt.Println("\nValues of Typed Constants are :-")
	fmt.Printf("Value of x is %v and x is of Type %T \n", x, x)
	fmt.Printf("Value of y is %v and y is of Type %T \n", y, y)
	fmt.Printf("Value of z is %v and z is of Type %T \n", z, z)
	fmt.Println("\nValues of UNTyped Constants are :-")
	fmt.Printf("Value of a is %v and a is of Type %T \n", a, a)
	fmt.Printf("Value of b is %v and b is of Type %T \n", b, b)
	fmt.Printf("Value of c is %v and c is of Type %T \n", c, c)
}
