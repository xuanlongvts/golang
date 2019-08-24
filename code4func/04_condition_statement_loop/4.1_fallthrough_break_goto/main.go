package  main

import (
	"fmt"
)

func main(){
	number := 10

	switch number {
	case 1:
		fmt.Println("number = 1")
		fallthrough
	case 10:
		goto hanldeNumberEqual10
		fmt.Println("number 1 = 10") // ignore this line
		hanldeNumberEqual10:
			fmt.Println("number 2 = 10")
	case 3:
		fmt.Println("number = 3")
		fallthrough
	case 4:
		fmt.Println("number = 4")
	}
}
