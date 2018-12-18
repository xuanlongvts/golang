package main

import "fmt"

func main() {
	s := struct {
		first   string
		friends map[string]int
		fav     []string
	}{
		first: "Long",
		friends: map[string]int{
			"Ronaldo": 7,
			"Messi":   10,
			"Neyma":   11,
		},
		fav: []string{
			"Music", "Football", "Code",
		},
	}

	fmt.Println("s=", s)
	for i, v := range s.friends {
		fmt.Println("i:", i, ", number:", v)
	}

	for _, v := range s.fav {
		fmt.Println("Hobbies: ", v)
	}
}
