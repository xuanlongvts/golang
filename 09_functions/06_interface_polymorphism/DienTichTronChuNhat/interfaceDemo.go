package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

func distance(x1, x2, y1, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.x2, r.y1, r.y2)
	w := distance(r.x1, r.x2, r.y1, r.y2)

	return l * w
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, v := range shapes {
		area += v.area()
	}

	return area
}

func main() {
	c := Circle{0, 0, 3}
	r := Rectangle{0, 0, 10, 10}

	areaC := totalArea(&c)
	areaR := totalArea(&r)
	fmt.Println("c=", areaC) // each area
	fmt.Println("r=", areaR) // each area

	totalT := totalArea(&c, &r)
	fmt.Println("total=", totalT) // total
}

/*
Trong các ví dụ trên, chúng ta đã định nghĩa 2 struct là Rectangle (hình chữ nhật) và Circle (hình tròn),
cả 2 struct này đều một phương thức tính diện tích có tên giống nhau là area().
Chúng ta có thể “gộp chung” 2 phương thức đó vào một kiểu dữ liệu khác có tên là Interface:

Trong ví dụ trên chúng ta định nghĩa một Interface có tên Shape, interface này có một phương thức là area().
Để định nghĩa một interface thì cũng giống như định nghĩa một struct, chúng ta dùng từ khóa type,
tiếp đến là tên interface rồi đến từ khóa interface, sau đó là danh sách các phương thức trong cặp dấu ngoặc nhọn {}.

Interface thực ra cũng không hẳn là một kiểu dữ liệu như struct vì interface chỉ chứa các phương thức chứ không chứa các trường,
interface cũng không có phần định nghĩa phương thức ở ngoài như các struct, chúng chỉ chứa tên phương thức là hết.
Vậy thì việc sử dụng interface có gì hay? Câu trả lời là chúng ta có thể dùng interface để thực hiện tính toán trên nhiều kiểu struct
khác nhau mà không quan tâm các struct đó là gì.

Trong ví dụ trên, chúng ta định nghĩa hàm totalArea() có chức năng tính tổng diện tích của bất cứ hình nào, hàm này nhận vào tham số là kiểu Shape,
nhưng chúng ta có thể truyền vào kiểu Circle hoặc kiểu Rectangle đều được, nếu chúng ta truyền vào kiểu Circle,
khi gọi phương thức area() thì Go sẽ gọi phương thức area() của struct Circle, và ngược lại khi truyền vào Rectangle thì gọi phương thức area()
của Rectangle.
*/
