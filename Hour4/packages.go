package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, This is my second program", 42, true)
	fmt.Println(n)
	fmt.Println(err)
}
