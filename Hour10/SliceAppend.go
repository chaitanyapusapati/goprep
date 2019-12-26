package main

import "fmt"

func main() {
	x := []int{2, 9, 23, 41}
	fmt.Println("Value of x is", x)
	fmt.Println("\nAppending Values to x ")
	x = append(x, 62, 79, 81)
	fmt.Println("Value of x is", x)
	y := []int{98, 112, 244, 398}
	fmt.Println("Value of y is", y)
	fmt.Println("\nAppending y Values to x ")
	x = append(x, y...)
	fmt.Println("Value of x is", x)
}
