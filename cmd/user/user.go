package user

import "fmt"

func Hello() {
	fmt.Println("hello")
}

type User struct {
	UserID    int64  `json:"user"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

var UserBase = []User{}
