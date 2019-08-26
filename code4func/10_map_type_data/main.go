package main

import "fmt"

func main() {

	// map[key data type] value data type
	// Zero value of map data is nil
	var mapDataOne map[string]int
	fmt.Println("map data one = ", mapDataOne)

	if mapDataOne == nil {
		fmt.Println("mapDataTwo = nil")
	}

	var mapDataTwo = make(map[string]int) // khai bao map thong qua make thi gia tri Zero value khong phai la nil
	fmt.Println("map data two = ", mapDataTwo)

	if mapDataTwo == nil {
		fmt.Println("mapDataTwo = nil")
	} else {
		fmt.Println("mapDataTwo !!!!!!!!!====== nil")
	}

	// declation and init
	var mapDataThree = map[string]int {
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}
	fmt.Println("mapDataThree = ", mapDataThree)

	// add new element, neu key trung thi se update gia tri sau cho gia tri truoc
	mapDataThree["key5"] = 5
	mapDataThree["key6"] = 6
	fmt.Println("mapDataThree = ", mapDataThree)

	// delete(mapData, key)
	delete(mapDataThree, "key6")
	fmt.Println("mapDataThree = ", mapDataThree)

	// length cua map
	fmt.Println("len of mapDataThree = ", len(mapDataThree))

	// map is a reference type data
	var mapDataFour = mapDataThree
	fmt.Println("=====================")
	fmt.Println("mapDataFour = ", mapDataFour)
	mapDataFour["key5"] = 555
	fmt.Println("mapDataFour = ", mapDataFour, ", mapDataThree = ", mapDataThree)

	// truy cap phan tu trong map
	var value, found = mapDataThree["key1000"]
	if found {
		fmt.Println("Tim thay gia tri tai key[1000] = ", value)
	} else {
		fmt.Println("Khong tim thay gia tri tai key = 1000")
	}

	// Trong map khong co toan tu == voi nhau (map chi co the so sanh voi nil)
	//if mapDataThree == mapDataFour {
	//	fmt.Println("map3 = map4")
	//}
}