package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: `barry`,
		last:  "allen",
	}
	p2 := person{
		first: `harrison`,
		last:  "wells",
	}
	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	fmt.Println("p1.first = ", p1.first)
	fmt.Println("p2.last = ", p2.last)
}
