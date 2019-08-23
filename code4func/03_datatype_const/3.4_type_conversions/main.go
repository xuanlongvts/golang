package  main

import "fmt"

func main(){
	var numberIntData int = 10
	var numberFloatData float32 = 1.1

	var total = float32(numberIntData) + numberFloatData
	fmt.Println("total = ", total)
}