package main

import "fmt"

type Ship struct {
	Engines int
	SOS     func()
}

func (s *Ship) Run() {
	fmt.Println("Run")
}

func DefaultSOS() {
	fmt.Println("SOS")
}

func main() {
	var s Ship

	s.SOS = DefaultSOS

	if s.Engines > 0 {
		s.Run() // Run est une méthode

	} else if s.SOS != nil {
		s.SOS() // SOS est un champ de type func()
	}
}
