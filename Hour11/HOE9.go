//Hands-on exercise #9
//Using the code from the previous example(HOE8.go), add a record to your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x[`flash`] = []string{`suit`, `speed`, `help`}
	for i, v := range x {
		fmt.Println("\nFor key ", i)
		for j, u := range v {
			fmt.Println("At Index", j, "value is", u)
		}
	}
}
