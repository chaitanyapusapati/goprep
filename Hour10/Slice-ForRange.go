package main

import "fmt"

func main() {
	x := []int{2, 9, 27, 42}
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
	fmt.Println("\nValue of x in Oth position ", x[0])
	fmt.Println("Value of x in 1st position ", x[1])
	fmt.Println("Value of x in 2nd position ", x[2])
	fmt.Println("Value of x in 3rd position ", x[3], "\n")
	for i, v := range x {
		fmt.Println("At Index", i, "Value in x is", v)
	}
}
