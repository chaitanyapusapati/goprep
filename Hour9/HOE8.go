// Hands-on exercise #8
// Create a program that uses a switch statement with no switch expression specified.
package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("TRUE, Prints it")
	case false:
		fmt.Println("FALSE, Not Prints it")
	default:
		fmt.Println("This is a Default Case")
	}
}
