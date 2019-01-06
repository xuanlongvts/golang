package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("noFile.txt")
	if err != nil {
		log.Panicln("Err happen", err)
	}
}

// Panicln is equivalent to Println() followed by a call panic()
