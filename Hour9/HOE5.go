// Hands-on exercise #5
// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println()
		fmt.Println(i, "/", 4, "=", i/4, "\t Rem =", i%4)
	}
}
