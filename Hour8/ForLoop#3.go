// Third Method of Writing FOR LOOP
package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}
}
