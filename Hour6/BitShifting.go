package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	a := 42
	fmt.Println("\nThe Value of a is ", a)
	fmt.Println("\nBit Shifting to Left side and Assigning to b")
	b := a << 1
	fmt.Printf("Value of %v after left shift in Decimal form is %v and in Binary form is %b\n", a, b, b)
	b = a >> 1
	fmt.Printf("Value of %v after Right shift in Decimal form is %v and in Binary form is %b\n", a, b, b)
	fmt.Printf("Variable b is of Type %T\n", b)
	fmt.Println("\nPerforming Left Bit Shifting by 10 times using Iota")
	fmt.Printf("Value of Kb (%v) after Bit Shifting in Binary form is %b\n", kb, kb)
	fmt.Printf("Value of mb (%v) after Bit Shifting in Binary form is %b\n", mb, mb)
	fmt.Printf("Value of gb (%v) after Bit Shifting in Binary form is %b\n", gb, gb)
	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
	)
	fmt.Println("\nPerforming Right Bit Shifting by 10 times")
	fmt.Printf("Value of Kb (%v) after Bit Shifting in Binary form is %b\n", kb, kb>>10)
	fmt.Printf("Value of mb (%v) after Bit Shifting in Binary form is %b\n", mb, mb>>10)
	fmt.Printf("Value of gb (%v) after Bit Shifting in Binary form is %b\n", gb, gb>>10)
}
