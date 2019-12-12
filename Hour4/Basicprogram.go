package main

import "fmt"

func main() {
	fmt.Println("Hello,vs code")
	fmt.Println("I am in main func")
	foo()
	fmt.Println("something more")
	boo()

}
func foo() {
	fmt.Println("Now I am in foo ")
}
func boo() {
	fmt.Println("And Now We Exicted")
}
