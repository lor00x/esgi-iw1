package main

import "fmt"

func main() {
	// Ecrivez une interface Role avec quelques méthodes
	// Ecrivez les structures concrètes qui implémentent l'interface
	// Initialisez un tableau/slice de Role
	// Appelez la même méthode de modification sur chacun des roles du tableau/slice
	// Appelez la même méthode d'affichage sur chacun des roles du tableau/slice
	// Ecrivez un switch/case qui affiche un message différent suivant
	// le type de chaque Rôle
	var roles = make([]Role, 0)
	user := User{
		Name:   "Harry",
		Active: true,
	}
	root := Root{}
	roles = append(roles, &user, &root)
	for _, role := range roles {
		role.SetActive(false)
		fmt.Println(role.GetName())
		SwitchCase(role)
	}
}

func SwitchCase(role Role) {
	switch role.(type) {
	case *User:
		fmt.Println("User")
	case *Root:
		fmt.Println("Root")
	default:
		fmt.Println("Unknow role")
	}
}
