package  main

import "fmt"

// https://mothereff.in/byte-counter  count byte of string
// http://www.fileformat.info/  find Unicode of character
/*
	1. Lý thuyết, mỗi ký tự chiếm 1 byte dữ liệu, mỗi khi lặp qua chữ Nhẫn sẽ in ra mỗi ký tự, tuy nhiên chữ Nhẫn có 6 byte dữ liệu.
	2. Kiểu rune không quan trọng mỗi ký tự chiếm bao nhiêu byte, mà tuỳ thuộc vào ký tự đó có bao nhiêu byte mà kiểu rune sẽ cấp phát bộ nhớ cho ký tự đó
	3. Bản chất của kiểu rune là alias của int32
*/

/*
	c: character
	U: Unicode
	x: decimal
*/
func main() {
	var stringShow string = "Nhẫn"
	fmt.Println("String show = ", stringShow) // Show result is right

	for i := 0; i < len(stringShow); i++ {
		fmt.Printf("%c", stringShow[i]) // Show result is not true
	}

	fmt.Println("\n==============================")
	var runeShow= []rune(stringShow)
	for j := 0; j < len(runeShow); j++ {
		fmt.Printf("%c", runeShow[j]) // show result is true
	}

	fmt.Println("\n==============================")
	var myRune= 'A'
	fmt.Printf("%T", myRune)

	fmt.Println("\n==============================")
	var runeMoreByte= '語' // 1 character but has 3 byte, find Unicode  8A9E	 http://www.fileformat.info/info/unicode/char/8a9e/index.htm
	fmt.Printf("%c -- %U", runeMoreByte, runeMoreByte)

	fmt.Println("\n==============================")
	var runeMoreByteOne= '🅐' // A character has 4 byte, find Unicode of A is 1F150
	fmt.Printf("%c -- %U", runeMoreByteOne, runeMoreByteOne)
}