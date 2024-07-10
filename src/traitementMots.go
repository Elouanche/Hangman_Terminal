package main

import "fmt"

//Vérifications du mots entré/S'il est deviné
func devinerMot(mot string, motJoueur string) bool {
	return motJoueur == mot
}
func traiterMotDevine(motAleatoire, choix string, tentatives *int) bool {
	if devinerMot(motAleatoire, choix) {
		clearTerminal()
		fmt.Println("Félicitations ! Vous avez deviné le mot.")
		return true
	} else {
		clearTerminal()
		fmt.Println("Raté, tu y étais presque, continue.")
		(*tentatives)--
		return false
	}
}
