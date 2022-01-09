package user

import (
	"fmt"
)

type User struct {
	Username string
	password int64
	email    string
}

func (u *User) UserLogin(data string) string {
	fmt.Println("data==", data)
	return "hello " + data
}
