/* Hands-on exercise #3
   Create a for loop using this syntax
   for condition { }
   Have it print out the years you have been alive. */
package main

import "fmt"

func main() {
	x := 1999
	for x <= 2019 {
		fmt.Println(x)
		x++
	}
}
