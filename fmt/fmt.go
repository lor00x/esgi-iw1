package main

import "fmt"

type City struct {
	Name    string
	Country string
}

func main() {
	c := City{
		Name:    "Paris",
		Country: "FR",
	}
	c2 := City{Name: "Paris", Country: "FR"}
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", c2)
}
