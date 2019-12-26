package main

import "fmt"

func main() {
	x := map[string]int{
		"goku":   1,
		"vegeta": 2,
	}
	fmt.Println("x is ", x)
	fmt.Println("Value for key goku is ", x["goku"])
	fmt.Println("\nChecking the Value and verifying for given key using v,ok")
	v, ok := x["gohan"]
	fmt.Println("Value for gohan", v)
	fmt.Println("Gohan is in x ", ok)
	fmt.Println("\nChecking the Value and verifying for given key using if condition")
	if v, ok = x["goku"]; ok {
		fmt.Println("Value for goku", v)
		fmt.Println("Goku is in x ", ok)
	}

}
