package main

import (
	"fmt"
	"time"
)

type item struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	CreatedBy      string `json:"created_by"`
	Approved       string `json:"approved"`
	ApprovedAt     string `json:"approved_at"`
	ApprovedBy     int    `json:"approved_by"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	MembersCount   int    `json:"members_count"`
	MembersList    []int  `json:"members_list"`
	OrderPins      string `json:"order_pins"`
	WorkGroupsUser struct {
		Id         int    `json:"id"`
		GroupId    int    `json:"group_id"`
		UserId     int    `json:"user_id"`
		Permission int    `json:"permission"`
		CreatedBy  int    `json:"created_by"`
		OrderPins  string `json:"order_pins"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	} `json:"work_groups_user"`
}

type T struct {
	OrderPins string `json:"order_pins"`
}

var data = []T{
	{
		OrderPins: "2022-02-01T14:51:02+07:00",
	},
	{
		OrderPins: "2022-02-01T15:01:43+07:00",
	},
	{
		OrderPins: "2022-02-01T14:50:50+07:00",
	},
	{
		OrderPins: "2023-08-20T14:41:21+07:00",
	},
}

func main() {
	//myDateString := "2022-02-01T14:50:50+07:00"
	//myDate, _ := time.Parse(time.RFC3339, myDateString)
	//miliSec := myDate.UnixMilli()
	//fmt.Println("miliSec: ", miliSec)

	for i := 0; i < len(data); i++ {
		for ii := i + 1; ii < len(data); ii++ {
			dateOne, _ := time.Parse(time.RFC3339, data[i].OrderPins)
			miliSecOne := dateOne.UnixMilli()

			dateTwo, _ := time.Parse(time.RFC3339, data[ii].OrderPins)
			miliSecTwo := dateTwo.UnixMilli()

			if miliSecOne < miliSecTwo {
				data[i], data[ii] = data[ii], data[i]
			}
		}
	}

	fmt.Println("data: ", data)
}
