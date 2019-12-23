/*Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.*/
package main

import "fmt"

func main() {
	x := 1999
	for {
		if x > 2019 {
			break
		}
		fmt.Println(x)
		x++
	}
}
