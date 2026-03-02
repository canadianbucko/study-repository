package main

import (
	"fmt"
)

type User struct {
	Name   string
	Rating float64
}

func (u *User) RatingUP(rating float64) {
	if u.Rating+rating <= 10 {
		u.Rating += rating
		fmt.Println("мы поменяли рейтинг")
	} else {
		fmt.Println("мы не меняели рейтинг")
		return
	}
}
func (u *User) ChangeNAME(name string) {
	if name == "" {
		fmt.Println("передано пустое имя")
		return
	} else {
		u.Name = name
	}
}

func main() {
	user1 := User{
		Name:   "кирюша",
		Rating: 5.2,
	}
	user1.ChangeNAME("кириешка")
	user1.RatingUP(4)

}
