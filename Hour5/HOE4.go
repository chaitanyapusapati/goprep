/* Hands-on exercise #4
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
In func main
Print out the value of the variable “x”
Print out the type of the variable “x”
Assign 42 to the VARIABLE “x” using the “=” OPERATOR
Print out the value of the variable “x” */
package main

import "fmt"

type own int

var x own

func main() {
	fmt.Println("\nThe Value of x is ", x)
	fmt.Printf("x is of Type %T \n", x)
	x = 42
	fmt.Println("\nAssigning Value to x ")
	fmt.Println("The Value of x is ", x)
}
