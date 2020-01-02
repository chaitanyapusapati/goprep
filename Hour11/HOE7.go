/*Hands-on exercise #7
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.*/
package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	z := [][]string{x, y}
	for i, v := range z {
		fmt.Println("\nAt Record ", i, "Values are")
		for j, u := range v {
			fmt.Println("In record", i, "at index", j, "Value is", u)
		}
	}
}
