package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
	}{
		first: "Harrison",
		last:  "wells",
	}
	fmt.Println("\np1 = ", p1)
	fmt.Println("p1.first = ", p1.first)
	fmt.Println("p1.last = ", p1.last)
}
