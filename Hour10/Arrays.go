package main

import "fmt"

func main() {
	var x [5]int
	var y [3]string
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println("\nAfter inserting Values")
	x[3] = 21
	y[0] = "bond,"
	y[1] = "JAMES"
	y[2] = "BOND"
	fmt.Println(x)
	fmt.Println(y)
}
