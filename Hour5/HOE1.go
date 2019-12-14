/*Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
		42
		“James Bond”
		true
Now print the values stored in those variables using
	-> a single print statement
	-> multiple print statements*/
package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("\nPrinting The Values In Single Print Statement\n")
	fmt.Println("The Value of x is ", x, "\tThe Value of y is ", y, "\tThe Value of z is ", z)
	fmt.Println("\nPrinting The Values In Multiple Print Statement")
	fmt.Println("The Value of x is ", x)
	fmt.Println("The Value of y is ", y)
	fmt.Println("The Value of z is ", z)
}
