package main

import "fmt"

func main() {
	l := []string{"Red", "yellow", "Blue", "silver", "white"}
	fmt.Println("Values in l are", l)
	s := []string{"yellow", "Red", "Black", "silver", "white"}
	fmt.Println("Values in s are", s)
	fmt.Println("\nCreating Multi-Dimensional Slices using l and s slice")
	msd := [][]string{l, s}
	fmt.Println(msd)
}
