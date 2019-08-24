package  main

import "fmt"

func Chao(){
	fmt.Println("Xin chao")
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func congrat(name string) string {
	result := fmt.Sprintf("Congrat %s", name)

	return result
}

// function return multi value
func rectInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// Named return value
func reactInfoOne(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main(){
	Chao()
	sayHello("Nong No")

	var congratResult = congrat("Nong No")
	fmt.Println(congratResult)

	var valOne, valTwo, area = rectInfo(10, 20)
	fmt.Println("valOne = ", valOne, ", valTwo = ", valTwo, ", area = ", area)

	fmt.Println("===================")
	var getWidth, getHeight, isSquare = reactInfoOne(100, 200)
	if(isSquare){
		fmt.Println("Square is true")
	} else {
		fmt.Println("width = ", getWidth, ", height = ", getHeight)
	}
}