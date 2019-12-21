package main

import "fmt"

func main() {
	s := "H"
	fmt.Printf("\ns is of Type %T\n", s)
	fmt.Println("The Value of s is ", s, "\n")
	fmt.Println("Converting to ByteString and Assigning to bs")
	bs := []byte(s)
	fmt.Printf("bs is of type %T\n", bs)
	fmt.Println("The Value of bs is", bs)
	fmt.Println("\nStoring the Index Zero(0) Value of bs in n(new variable)")
	n := bs[0]
	fmt.Printf("n is of Type %T\n", n)
	fmt.Println("The Value of n is", n)
	fmt.Println("\nConverting the Decimal Value of n into other Numerical Types")
	fmt.Printf("Value of n in Binary is %b\n", n)
	fmt.Printf("Value of n in Octal  is %#o\n", n)
	fmt.Printf("Value of n in Hexa Decimal is %#X\n", n)
}
