package main

import "fmt"

func main() {
	x := []int{2, 9, 21, 91, 245, 962, 4647, 84752}
	fmt.Println("Values in x ", x)
	fmt.Println("\nDeleting 21,91 from  x ")
	x = append(x[:1], x[4:]...)
	fmt.Println("Values in x ", x)
}
