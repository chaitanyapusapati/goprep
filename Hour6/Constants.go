package main

import "fmt"

const (
	a = 42
	b = 42.78
	c = "James Bond"
)

func main() {
	fmt.Println("\nThe Value of a is ", a)
	fmt.Println("The Value of b is ", b)
	fmt.Println("The Value of c is ", c)
	fmt.Println("\nLets See The Types Of these Variables")
	fmt.Printf("Variable a is of Type %T\n", a)
	fmt.Printf("Variable b is of Type %T\n", b)
	fmt.Printf("Variable c Is of Type %T\n", c)
	/* This Doesn't Work Because a,b,c are constants and we cannot change the value of constants once they are Assigned
	a = 4683
	b = 47.4455589758
	c = "Bond"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) */
}
