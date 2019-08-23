package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("noFile.txt")
	if err != nil {
		fmt.Println("fmt Err happen:", err)
		log.Println("log Err happen:", err)
	}
}
