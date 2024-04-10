package main

type User struct {
	Name string
	Active bool
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetActive(active bool) {
	u.Active = active
}

func (u *User) IsAdmin() bool {
	return false
}
