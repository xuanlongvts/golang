package main

import (
	"fmt"
	"math"
)

// Giả sử chúng ta có đoạn code tính diện tích hình tròn và diện tích hình chữ nhật từ các điểm trên mặt phẳng như sau:
func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, x2, y1, y2 float64) float64 {
	l := distance(x1, x2, y1, y2)
	w := distance(x1, x2, y1, y2)

	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
}

/*
	Thoạt nhìn thì bài toán này có vẻ đơn giản, tuy nhiên khi bạn muốn tính diện tích của khoảng vài chục hình vuông/hình tròn,
	lúc này bản thân việc đặt tên biến đã muốn mệt chứ chưa nói đến việc tính toán, dĩ nhiên bạn có thể dùng array, slice…
	nhưng giả sử bạn muốn hình dung trong đầu hình nào ở vị trí số mấy trong mảng là cũng khá khó khăn rồi.
	Do đó bạn cần dùng kiểu struct để có thể quản lý tốt hơn.
*/
