//Hands-on exercise #2
//Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.
package main

import "fmt"

type person struct {
	first        string
	last         string
	ice_flavours []string
}

func main() {
	p1 := person{
		first:        "harr",
		last:         "wells",
		ice_flavours: []string{"vanilla", "butterscoth", "chocolate"},
	}
	p2 := person{
		first:        "cisco",
		last:         "ramon",
		ice_flavours: []string{"vanilla", "strawberry", "butterscotch"},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for i, v := range m {
		fmt.Println("\nFor key ", i, "The Values are")
		fmt.Println(v.first)
		fmt.Println(v.last)
		for j, u := range v.ice_flavours {
			fmt.Println(j, u)
		}
	}

	fmt.Println("\np1.first =", p1.first)
	fmt.Println("p1.last =", p1.last)
	for i, v := range p1.ice_flavours {
		fmt.Println(i, v)
	}
	fmt.Println("\np2.first =", p2.first)
	fmt.Println("p2.last =", p2.last)
	for i, v := range p2.ice_flavours {
		fmt.Println(i, v)
	}

}
