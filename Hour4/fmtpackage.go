package main

import "fmt"

var y = 42

func main() {
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("Binary value of y is %b\n", y)
	fmt.Printf("HexaDeccimal value of y  is %x\n", y)
	fmt.Printf("HexaDecimal value of y with 0x is %#x\n", y)
	fmt.Printf("Octal value of y is %o\n", y)
	fmt.Printf("Octal value of y with 0x is %#o\n", y)
	fmt.Printf("Value of y is %v\n", y)
	fmt.Printf("Decimal value of y is %d\n", y)
	fmt.Println("Using Sprintf as input for s")
	s := fmt.Sprintf("%T\t%b\t%x\t%o\n", y, y, y, y)
	fmt.Printf(s)
}
