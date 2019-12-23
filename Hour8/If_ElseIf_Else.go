package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("Value of x is 40")
	} else if x == 41 {
		fmt.Println("Value of x is 41")
	} else if x < 41 {
		fmt.Println("Value of x is less than 41")
	} else {
		fmt.Println("Value of x is greater than 41")
	}
}
