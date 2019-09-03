package  main

import "fmt"

func main(){
	 var arrA = [5]int {1, 2, 3, 4, 5}
	 var sliceOne = arrA[:]
	 fmt.Println("slice one before = ", sliceOne)

	 sliceOne[1] = 777
	 fmt.Println("slice one after = ", sliceOne)
	 fmt.Println("arrA = ", arrA)

	 // length and capacity
	 // 1. length: So luong phan tu cua slice
	 // 2. capacity: So luong phan tu underlying array bat dau vi tri start khi ma slice duoc tao
	 var sliceCountries = [...]string {"VN", "USA", "CANADA", "FRANCE", "IRan", "Lao", "Campuchia"}
	 sliceTwo := sliceCountries[2:5]
	 fmt.Println("sliceTwo: = ", sliceTwo, ", length = ", len(sliceTwo), ", capacity = ", cap(sliceTwo))

	 fmt.Println("==================================")
	 // make, append and copy

	 // 1. make(kieu du lieu, length, capacity) neu thieu capacity thi capaicity = length
	 var sliceThree = make([]int, 2, 5) //
	 fmt.Println("Slice three = ", sliceThree, ", length = ", len(sliceThree), ", capacity = ", cap(sliceThree))

	var sliceFour = make([]int, 2) //
	fmt.Println("Slice Four = ", sliceFour, ", length = ", len(sliceFour), ", capacity = ", cap(sliceFour))

	fmt.Println("==================================")
	// 2. append
	var sliceFive []int
	sliceFive = append(sliceFive, 100)
	fmt.Println("Sclice five = ", sliceFive)

	var sliceSix = make([]int, 2)
	sliceSix = append(sliceSix, 999)
	fmt.Println("Sclice six = ", sliceSix)

	fmt.Println("==================================")
	// 3. copy
	src := []string {"A", "B", "C", "D"}
	dest := make([]string, 2)
	destTwo := make([]string, 3) // length will number element copy
	number := copy(dest, src)
	number2 := copy(destTwo, src)
	fmt.Println("number = ", number, ", dest = ", dest)
	fmt.Println("number 2 = ", number2, ", dest two = ", destTwo)

	fmt.Println("==================================")
	// 4. delete (golang not give api but use append is a trick)
	var newSrc = append(src[:1], src[2:]...)
	fmt.Println("new src = ", newSrc)
}