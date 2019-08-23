package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{
		first: "Long",
		last:  "Le",
		fav:   []string{"music", "football", "cafe"},
	}
	p2 := person{
		first: "Nong",
		last:  "No",
		fav:   []string{"soccer", "basketball", "volleyball"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println("m = ", m)
	for k, v := range m {
		fmt.Println("k=", k)
		fmt.Printf("\t v= %v\n", v.first)
		for _, val := range v.fav {
			fmt.Printf("\t val: %v\n", val)
		}
	}
}
