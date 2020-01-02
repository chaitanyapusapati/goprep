/*Hands-on exercise #3
start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Use SLICING to create the following new slices which are then printed:
	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47] */
package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("\nUsing Slicing to Creating new Slices from x")
	y := x[:5]
	z := x[5:]
	a := x[2:7]
	b := x[1:6]
	fmt.Println("Slice y is", y)
	fmt.Println("Slice z is", z)
	fmt.Println("Slice a is", a)
	fmt.Println("Slice b is", b)
}
