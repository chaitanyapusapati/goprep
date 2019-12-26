package main

import "fmt"

func main() {
	x := []int{2, 9, 27, 42, 59}
	fmt.Println("Value of x ", x)
	fmt.Println("\nSlicing operation")
	fmt.Println("Slicing with empty values :-", x[:])
	fmt.Println("Slicing from index 1 to 4 :-", x[1:4])
	fmt.Println("Slicing from index 0 to 3 :-", x[0:3])
}
