package main

import (
	"fmt"
)

var y = 42
var z = "shaken,not stirred"
var a string = `Todd said"GO SUFFERS,NO FOOLS



yes
he
said
that"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
