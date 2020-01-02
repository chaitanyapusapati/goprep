//Hands-on exercise #10
//Using the code from the previous example(HOE9.go), delete a record from your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	x[`flash`] = []string{`suit`, `speed`, `help`}
	delete(x, `moneypenny_miss`)
	for i, v := range x {
		fmt.Println("\nFor key ", i)
		for j, u := range v {
			fmt.Println("At Index", j, "value is", u)
		}
	}
}
