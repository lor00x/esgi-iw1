package main

type Role interface {
	GetName() string
	SetActive(bool)
	IsAdmin() bool
}
