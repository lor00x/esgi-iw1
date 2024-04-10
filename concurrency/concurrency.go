package main

// Goroutine:
// Clients: en boucle
//      - passent commande au serveur
//		- boivent
//		- attendent
// Serveur:
//		- reçoit des commandes de boisson
//      - prépare la boisson et la renvoie au client
// Patron:
//		- attends 5 secondes
//		- prévient tout le monde d'arrêter
//
// Hints:
// - Lancer une goroutine: go func()
// - Créer un channel: make(chan MyType)
