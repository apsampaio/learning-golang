package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	login := "password1234"

	err = bcrypt.CompareHashAndPassword(bs, []byte(login))

	if err != nil {
		fmt.Println("LOGIN FAILED, PASSWORD WRONG")
		return
	}

	fmt.Println("YOU LOGGED IN")

}
