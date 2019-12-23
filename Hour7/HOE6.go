/* Hands-on exercise #6
   Using iota, create 4 constants for the NEXT 4 years. Print the constant values. */
package main

import "fmt"

const (
	x = 2019 + iota
	y = 2019 + iota
	z = 2019 + iota
	a = 2019 + iota
)

func main() {
	fmt.Println("\ncreating Constants for NEXT 4 years")
	fmt.Println("Value of Constant x is ", x)
	fmt.Println("Value of Constant y is ", y)
	fmt.Println("Value of Constant z is ", z)
	fmt.Println("Value of Constant a is ", a)
}
