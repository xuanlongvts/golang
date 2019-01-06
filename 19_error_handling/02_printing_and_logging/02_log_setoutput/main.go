package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f1, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Err1 ", err)
	}
	defer f1.Close()
	log.SetOutput(f1)

	f2, err := os.Open("noFile.txt")
	if err != nil {
		log.Println("Err2 happen: ", err)
	}
	defer f2.Close()

	fmt.Println("Check the log.txt file in the directory")
}
