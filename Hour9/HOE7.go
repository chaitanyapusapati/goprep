// Hands-on exercise #7
// create a program that uses "if" , “else if” and “else”.
package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("x is Equal to 40")
	} else if x < 40 {
		fmt.Println("x is Less than 40")
	} else {
		fmt.Println("x is Greater than 40")
	}
}
