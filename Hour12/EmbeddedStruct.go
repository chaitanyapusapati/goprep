package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretagent struct {
	person
	per bool
}

func main() {
	s1 := secretagent{
		person: person{
			first: "sherlock",
			last:  "wells",
		},
		per: true,
	}
	fmt.Println("s1 = ", s1)
	fmt.Println("s1.first = ", s1.first)
	fmt.Println("s1.last = ", s1.last)
	fmt.Println("s1.per = ", s1.per)

}
