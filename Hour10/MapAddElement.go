package main

import "fmt"

func main() {
	x := map[string]int{
		"goku":  1,
		"gohan": 2,
	}
	fmt.Println("x is", x)
	fmt.Println("\nAdding a Element in Slice")
	x["goten"] = 3
	for i, v := range x {
		fmt.Println("The key is", i, "and Value is ", v)
	}
}
