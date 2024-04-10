package main

import (
	"fmt"
	"io"
)

type User struct {
	Name string
}

func (u *User) SetName(n string) {
	u.Name = n
}

func NewUser() *User {
	u := User{}
	u.SetName("Harry") // OK
	fmt.Println("Name:", u.Name)
	return &u
	io.ReadCloser()
}
