package main

import "fmt"

func main() {
	for i := 33; i < 123; i++ {
		fmt.Printf("Decimal is %d\t HexaDecimal is %#X\tUNICODE is %#U\n\n", i, i, i)
	}
}
