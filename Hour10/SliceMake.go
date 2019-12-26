package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
	x[9] = 21
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
	x = append(x, 42)
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
	x = append(x, 47)
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
	x = append(x, 53)
	fmt.Println("\nValue of x is ", x)
	fmt.Println("Length of x is ", len(x))
	fmt.Println("Capacity of x is ", cap(x))
}
