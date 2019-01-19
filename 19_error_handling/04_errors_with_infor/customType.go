package main

import (
	"fmt"
	"log"
	"math"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("A norgate math error occured: %v %v %v ", n.lat, n.long, n.err)
}

func main() {
	const number = -64
	val, err := sqrt(number)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("sqrt of %v is %v \n", number, val)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("Norgate math redux: square root of negative number: %v", f)

		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}

	return math.Sqrt(f), nil
}
