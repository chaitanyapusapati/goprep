/*Hands-on exercise #5
Building on the code from the previous excersice
at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
eg:
		type own int
		var x own
		var y int

In func main
This should already be done
	->print out the value of the variable “x”
	->print out the type of the variable “x”
	->assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
	->print out the value of the variable “x”
Now do this
Now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
then use the “=” operator to ASSIGN that value to “y”
	->print the value stored in “y”
	->print the type of “y” */
package main

import "fmt"

type own int

var x own
var y int

func main() {
	fmt.Println("\nThe Value of x is ", x)
	fmt.Printf("x is of Type %T \n", x)
	x = 42
	fmt.Println("\nAssigning Value to x ")
	fmt.Println("The Value of x is ", x)
	fmt.Println("\nConverting x to its Underlying Type and Assigning to y")
	y = int(x)
	fmt.Println("The Value of y is ", y)
	fmt.Printf("y is of Type %T \n", y)
}
