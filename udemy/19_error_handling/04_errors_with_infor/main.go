package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	const number = -81
	val, err := sqrt(number)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("sqrt of %v is %v \n", number, val)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Norgate math: square of negative number")
		// return 0, fmt.Errorf("Norgate math: square of negative number: %v", f)
	}

	return math.Sqrt(f), nil
}
