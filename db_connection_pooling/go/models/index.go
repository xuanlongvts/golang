package models

type Products struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type Response struct {
	Number   int64       `json:"number"`
	Elapsed  int64       `json:"elapsed"`
	Average  float64     `json:"average"`
	Products []*Products `json:"products"`
}
