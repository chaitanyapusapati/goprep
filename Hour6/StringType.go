package main

import "fmt"

func main() {
	s := "Hey, Strings"
	fmt.Printf("s is of type %T \n", s)
	fmt.Println("The Value of s is ", s)
	fmt.Println("\nConverting The String to Bytes and Storing in bs")
	bs := []byte(s)
	fmt.Printf("bs is of Type %T \n", bs)
	fmt.Println("The Value of bs is ", bs, "\n")
	for v, i := range s {
		fmt.Printf("For the Char %c at index %v the ByteString is %v and the Utf code is %#U \n", i, v, i, s[v])
	}
}
