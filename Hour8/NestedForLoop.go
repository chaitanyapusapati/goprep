package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("\nValue of i in Outer Loop is ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("\t\tValue of j in Inner loop is ", j)
		}
	}
}
