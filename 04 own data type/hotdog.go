package main

import (
	"fmt"
)

var y int

type hotdog int

var b hotdog

func main() {
	y = 42
	b = 43
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
