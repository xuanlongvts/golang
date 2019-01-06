package main

import (
	"os"
)

func main() {
	_, err := os.Open("noFile.txt")
	if err != nil {
		panic(err)
	}
}

// the Panic function call panic after the log message
