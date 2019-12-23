package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("TRUE, Prints it")
		fallthrough
	case false:
		fmt.Println("FALSE, Not Prints it")
		fallthrough
	case 2 == 2:
		fmt.Println("TRUE, Prints it")
		fallthrough
	case 2 != 2:
		fmt.Println("FALSE, Not Prints it")
		fallthrough
	default:
		fmt.Println("This is Default Case")
	}
}
