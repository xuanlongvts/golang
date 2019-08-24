package main

import "fmt"

func main()  {
	countries := [...]string {"VN", "US", "Canada"}

	// cach 1
	for i := 0; i < len(countries); i++ {
		fmt.Print(countries[i], ", ")
	}

	// cach 2
	for index, value := range countries {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// blank identifier
	fmt.Println("===================")
	for _, val := range countries {
		fmt.Println("value = ", val)
	}

	// mang hai chieu [row][column]
	fmt.Println("===================")
	arrNumbers := [4][2]int {
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println("arrNumbers = ", arrNumbers)
	fmt.Println("===================")
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(arrNumbers[i][j], "  ")
		}

		fmt.Println("")
	}
}