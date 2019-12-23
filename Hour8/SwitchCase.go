package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Println("\nEnter a Suit Colour\nR - Red \t Y - Yellow \t B - Black \nS-Silver \t W-White \t O-Other colour")
	fmt.Scanln(&x)
	x = strings.ToUpper(x)
	fmt.Printf("YOU ARE ")
	switch x {
	case "R":
		fmt.Println("FLASH")
	case "Y":
		fmt.Println("REVERSE FLASH")
	case "B":
		fmt.Println("ZOOM")
	case "S":
		fmt.Println("SAVITAR")
	case "W":
		fmt.Println("GOD SPEED")
	default:
		fmt.Println("Rookie Speedster")
	}
}
