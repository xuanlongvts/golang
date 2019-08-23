package main

import (
	"fmt"
	"math"
	"math/bits"
)

/*
	https://www.eso.org/~ndelmott/ascii.html (ASCII Character Chart with Decimal, Binary and Hexadecimal Conversions)

	1. bool
	2. string
	3. int int8 int16 int32 int64 (so nguyen)
		3.1 uint uint8 uint16 uint32 uint64 (so nguyen duong)
	4. byte (alias for uint8)
	5. rune (alias for uint32) represents a Unicode code point
	6. float32, float64 (so thuc)
	7. complex64, complex128 (so phuc)
*/

func main() {
	var bool1 bool = true
	fmt.Println("bool1 = ", bool1)

	var string1 string = "aaa"
	fmt.Println("string show = ", string1)

	// int
	var intShow int = 123
	fmt.Println("int show = ", intShow)

	// int 8, 16, 32, 64
	// 1. Range
	// 2. Bits

	// 1. Range
	// ======================================
	// int 8	-128 -> 127
	fmt.Println("min int8 = ", math.MinInt8, ", max int8 = ", math.MaxInt8)

	// int 16	-32768 -> 32767
	fmt.Println("min int16 = ", math.MinInt16, ", max int16 = ", math.MaxInt16)

	// int 32	-2147483648 -> 2147483647
	fmt.Println("min int32 = ", math.MinInt32, ", max int32 = ", math.MaxInt32)

	// int 64	-9223372036854775808 -> 9223372036854775807
	fmt.Println("min int64 = ", math.MinInt64, ", max int64 = ", math.MaxInt64)

	// 2. Bits
	// ======================================
	fmt.Println("Bits: ======================================")
	fmt.Println("Amount bit: int8 = ", bits.OnesCount8(math.MaxUint8), ", " + "int16 = ", bits.OnesCount16(math.MaxUint16), ", int32 = ", bits.OnesCount32(math.MaxUint32), ", " +
		"int64 = ", bits.OnesCount64(math.MaxUint64))

	// uint
	fmt.Println("Uint: ======================================")
	var myUint uint = 1
	fmt.Println("uint = ", myUint)

	// byte
	fmt.Println("byte: ======================================")
	var myByte byte = 255 // byte = 256 will show error because range of int8 is from -128 -> 127
	fmt.Printf("value myByte = %v", myByte)
	fmt.Printf("\ndata Type = %T", myByte)

	fmt.Println("")
	var myByteOne byte = 'A'
	fmt.Println("Acscci of A character = ", myByteOne)

	var myByteTwo byte = 'D'
	fmt.Printf("Hexa of D character = %X", myByteTwo)
}