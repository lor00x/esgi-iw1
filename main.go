package main

import "fmt"

func main() {
	msg := "Happy birthday ! 🎉🥳🎂"

	r := []byte(msg)

	msg2 := string(r)

	fmt.Println(msg2)
}
