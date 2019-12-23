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
	a := (42 == 43)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	fmt.Println("\nThe Value of a is ", a)
	fmt.Println("The Value of b is ", b)
	fmt.Println("The Value of c is ", c)
	fmt.Println("The Value of d is ", d)
	fmt.Println("The Value of e is ", e)
	fmt.Println("The Value of f is ", f)

}
