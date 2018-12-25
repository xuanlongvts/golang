package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "123456"
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println("errMess: ", err)
	}
	fmt.Println("pass =", pass)
	fmt.Println("bs =", bs)

	loginPass := "123456 "
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))
	if err != nil {
		fmt.Println("You can't login: ", err)
		return
	}

	fmt.Println("Login success")
}
