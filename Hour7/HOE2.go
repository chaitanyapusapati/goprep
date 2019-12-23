/* Hands-on exercise #2
   Using the following operators, write expressions and assign their values to variables:
	==
	<=
	>=
	!=
	<
	>
Now print each of the variables. */
package main

import "fmt"

func main() {
	x := 0
	y := 0
	fmt.Println("\nEnter the values for  variable x")
	fmt.Scanln(&x)
	fmt.Println("Enter the values for  variable y")
	fmt.Scanln(&y)
	fmt.Println("The Values you have entered are x = ", x, ", y = ", y)
	if x == y {
		x = x - 0
	}
	if x <= y {
		x = x + 1
	}
	if x >= y {
		x -= 2
	}
	if x != y {
		y += 3
	}
	if x < y {
		y -= 4
	}
	if x > y {
		x -= 5
	}
	fmt.Println("\nAfter Executing the Expressions and Assigning the values to the given Variables")
	fmt.Println("The Value of x is ", x)
	fmt.Println("The Value of y is ", y)
}
