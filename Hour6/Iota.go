package main

import "fmt"

const (
	a = iota
	b
	c
	g
	h
	i = iota + 1
)
const (
	d = iota
	e
	f
)

func main() {
	fmt.Println("\nThe Value of a is ", a)
	fmt.Println("The Value of b is ", b)
	fmt.Println("The Value of c is ", c)
	fmt.Println("The Value of g is ", g)
	fmt.Println("The Value of h is ", h)
	fmt.Println("The Value of i is ", i, "\n")
	fmt.Println("The Value of d is ", d)
	fmt.Println("The Value of e is ", e)
	fmt.Println("The Value of f is ", f)
}
