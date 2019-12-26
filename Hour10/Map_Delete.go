package main

import "fmt"

func main() {
	x := map[string]int{
		"goku":  1,
		"gohan": 2,
	}
	fmt.Println("values in x are", x)
	fmt.Println("\nDeleting goku")
	delete(x, "goku")
	fmt.Println("values in x are", x)
	if v, ok := x["gohan"]; ok {
		fmt.Println("\nValue of gohan is ", v)
		fmt.Println("gohan is in x ", ok)
		fmt.Println("\nDeleting gohan")
		delete(x, "gohan")
	}
	fmt.Println("values in x are", x)
}
