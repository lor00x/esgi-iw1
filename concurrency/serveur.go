package main

import "fmt"

type Commande struct {
	Boisson string
	Client  chan string
}

func Serveur(commandes <-chan Commande, patron <-chan bool, ok chan<- bool) {
	for {
		select {
		case commande := <-commandes:
			fmt.Println("Serveur, j'ai reçu une commande de:", commande.Boisson)
			boisson := commande.Boisson
			commande.Client <- boisson
		case <-patron:
			fmt.Println("Serveur: j'ai fini mon travail")
			ok <- true
			return
		}
	}
}
