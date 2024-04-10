package main

import (
	"fmt"
	"time"
)

func Client(serveur chan<- Commande, patron <-chan bool, ok chan<- bool) {
	for {
		select {
		case <-patron:
			fmt.Println("J'ai bien bu !")
			ok <- true
			return
		default:
			client := make(chan string)
			commande := Commande{
				Boisson: "Mojito",
				Client:  client,
			}
			fmt.Println("Je commande !")
			serveur <- commande
			boisson := <-commande.Client
			fmt.Println("J'ai reçu une boisson ", boisson)
			time.Sleep(time.Second)
		}
	}
}
