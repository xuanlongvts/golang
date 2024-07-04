package main

type caculate interface {
	Plus(a int, b int) int
	Minus(a int, b int) int
}

type dataReq struct {
	A int `json:"a"`
	B int `json:"b"`
}

func (req *dataReq) Plus(a int, b int) int {
	return a + b
}

func (req *dataReq) Minus(a int, b int) int {
	return a - b
}
