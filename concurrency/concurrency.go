package main

import (
	"fmt"
	"time"
)

// Goroutine:
// Clients: en boucle
//   - passent commande au serveur
//   - boivent
//   - attendent
//
// Serveur:
//   - reçoit des commandes de boisson
//   - prépare la boisson et la renvoie au client
//
// Patron:
//   - attends 5 secondes
//   - prévient tout le monde d'arrêter
//
// Hints:
// - Lancer une goroutine: go func()
// - Créer un channel: make(chan MyType)
func main() {
	commandes := make(chan Commande)
	doneServeur := make(chan bool)
	doneClient := make(chan bool)
	okServeur := make(chan bool)
	okClient := make(chan bool)
	go Serveur(commandes, doneServeur, okServeur)
	for range 5 {
		go Client(commandes, doneClient, okClient)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("C'est terminé !")

	for range 5 {
		doneClient <- true
	}

	fmt.Println("J'attends les client")
	for range 5 {
		<-okClient
	}

	doneServeur <- true
	<-okServeur

	fmt.Println("A demain!")

}
