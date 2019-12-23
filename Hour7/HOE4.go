/* Hands-on exercise #4
Write a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in Decimal, Binary, and Hex */
package main

import "fmt"

func main() {
	var x int
	fmt.Println("\nEnter an Integer")
	fmt.Scanln(&x)
	fmt.Printf("Value of %v in Decimal Format is %d \n", x, x)
	fmt.Printf("Value of %v in Binary Format is %b \n", x, x)
	fmt.Printf("Value of %v in Hex Format is %#X \n", x, x)
	fmt.Println("\nPerformimg the Left Bit Shifting and Assigning to Variable to y")
	y := x << 1
	fmt.Printf("Value of %v in Decimal Format is %d \n", y, y)
	fmt.Printf("Value of %v in Binary Format is %b \n", y, y)
	fmt.Printf("Value of %v in Hex Format is %#X \n", y, y)
}
