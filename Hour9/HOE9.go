// Hands-on exercise #9
// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "football":
		fmt.Println("Football is Favourite Sport")
	case "BadMinton":
		fmt.Println("BadMinton is Favourite Sport")
	case "cricket":
		fmt.Println("Cricket is Favourite Sport")
	default:
		fmt.Println("I got NO Favourite Sport")
	}
}
