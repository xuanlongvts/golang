package main

import (
	"fmt"
)

func main() {
	favSport := "football"

	switch favSport {
	case "football":
		fmt.Println("Go to the football pitches")
		fallthrough // move to the next case
	case "skiing":
		fmt.Println("Go to the mountain")
	case "swimming":
		fmt.Println("Go to the pool")
	case "furfing":
		fmt.Println("Go to the hawaii")
	case "wingsuit flying":
		fmt.Println("What would you like me to say at your funeral")
	}
}
