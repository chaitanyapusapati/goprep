/* Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors. */
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
