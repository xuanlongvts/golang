package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("noFile.txt")
	if err != nil {
		log.Fatalln("Err happen: ", err)
	}

	// foo()
}

func foo() {
	fmt.Println("When os.Exit() is called, defer function don't run")
}

// Fatal function call os.Exit(1) after writting the log message
// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
