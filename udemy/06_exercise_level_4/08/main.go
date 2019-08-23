package main

import "fmt"

func main() {
	m := map[string][]string{
		`Le`:   []string{`Mot`, `Hai`, `Ba`},
		`Xuan`: []string{`Bon`, `Nam`, `Sau`},
		`Long`: []string{`Bay`, `Tam`, `Chin`},
	}
	fmt.Println("m = ", m)

	m[`Nong`] = []string{`Muoi`, `Muoi Mot`, `Muoi Hai`} // add new element
	for i, v := range m {
		fmt.Println("record: ", i)
		for ii, vv := range v {
			fmt.Printf("\t index: %v , value = %v \n", ii, vv)
		}
	}
	fmt.Println("===========================================")
	delete(m, `Nong`) // delete an item
	for i, v := range m {
		fmt.Println("record: ", i)
		for ii, vv := range v {
			fmt.Printf("\t index: %v , value = %v \n", ii, vv)
		}
	}
}
