package main

import "fmt"

var x, y int = 21,21
var z string

func main() {
	fmt.Println(x)
	fmt.Println(1 + y)
	z= "shaken!, Not Stirred"
	fmt.Println(z)
	a:=`This takes the raw string like :
			"Shaken, Not Stirred"
						cause we are using back quotes`
	fmt.Println(a)

}
