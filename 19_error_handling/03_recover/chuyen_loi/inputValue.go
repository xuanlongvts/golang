package main

import "fmt"

func main() {

	var n int
	fmt.Print("Nhap so nguyen duong: ")

	// _, err := fmt.Scan(&n)
	_, err := fmt.Scanf("%d", &n)
	if err == nil && n < 0 {
		fmt.Printf("\t Loi, %d la so nguyen am\n", n)
	}

	fmt.Println("Ban vua nhap so: ", n)
}
