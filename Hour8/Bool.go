package main

import "fmt"

func main() {
	if true {
		fmt.Println(1)
	}
	if false {
		fmt.Println(2)
	}
	if !true {
		fmt.Println(3)
	}
	if !false {
		fmt.Println(4)
	}
	if 2 == 2 {
		fmt.Println(5)
	}
	if 2 != 2 {
		fmt.Println(6)
	}
	if !(2 == 2) {
		fmt.Println(7)
	}
	if !(2 != 2) {
		fmt.Println(8)
	}
}
