package main

import "fmt"

func main() {
	fmt.Println("\nUsing CONTINUE Keyword to print the Even Numbers till 100")
	for i := 0; i < 101; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
