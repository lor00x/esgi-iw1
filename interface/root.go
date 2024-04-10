package main

type Root struct {
}

func (u *Root) GetName() string {
	return "root"
}

func (u *Root) SetActive(active bool) {}

func (u *Root) IsAdmin() bool {
	return true
}
