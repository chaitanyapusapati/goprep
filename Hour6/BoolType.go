package main

import "fmt"

var x bool

func main() {
	fmt.Println("\nThe Value of x is ", x)
	x = true
	fmt.Println("The Value of x is ", x)
	a := 42
	b := 21
	fmt.Println("The Value of a and b are ", a, ",", b)
	fmt.Println("\nIn Go We can Print Bool Value as an Output")
	fmt.Println("a==b ", a == b)
	fmt.Println("a>b ", a > b)
	fmt.Println("a>=b ", a >= b)
	fmt.Println("a<=b ", a <= b)
	fmt.Println("a<b ", a < b)
	fmt.Println("a!=b ", a != b)
}
