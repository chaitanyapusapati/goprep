// Write a program that prints a number in decimal, binary, and hex
package main

import "fmt"

func main() {
	var n int
	fmt.Printf("\nEnter a Number :- ")
	fmt.Scanln(&n)
	fmt.Println("\nPrinting the Number in Numerical Formats")
	fmt.Printf("Number %v in Decimal Format is %d \n", n, n)
	fmt.Printf("Number %v in Binary Format is %b \n", n, n)
	fmt.Printf("Number %v in octal Format is %#o \n", n, n)
	fmt.Printf("Number %v in Hexa-Decimal Format is %#X \n", n, n)
}
