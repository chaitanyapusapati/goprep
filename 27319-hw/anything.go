package main

import "fmt"

func main() {
	fmt.Println("i am in main")
	foo()
	fmt.Println("out of foo and in main")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
	}


func foo() {
	fmt.Println("i am in foo")
}
func bar(){
	fmt.Println("and then we exited")
}