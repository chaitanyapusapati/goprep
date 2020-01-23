//Hands-on exercise #4
//Create and use an anonymous struct.
package main

import "fmt"

func main() {
	s := struct {
		first        string
		friends      map[string][]string
		ice_flavours []string
	}{
		first: "Dr.wells",
		friends: map[string][]string{
			"wells":  {"Barry", "cisco", "caitlin"},
			"icicle": {"Dr.wells", "caitlin", "Barry"},
		},
		ice_flavours: []string{"vanilla", "chocolate", "ButterSctoch"},
	}
	fmt.Println(s.first)
	for i, v := range s.friends {
		fmt.Println("\nFriends of ", i, "are")
		for j, u := range v {
			fmt.Println(j, u)
		}
	}
	fmt.Println("\nFavourte Ice cream Flavours of Dr.wells are")
	for i, v := range s.ice_flavours {
		fmt.Println(i, v)
	}
}
