package main

import "fmt"

//Texte du jeu
func afficherEtat(motMasque string, tentatives int) {
	fmt.Printf("Mot masqué: %s\n", motMasque)
	fmt.Printf("Il vous reste %d tentatives.\n", tentatives)
	fmt.Print("Choisissez une lettre ou devinez le mot: ")
}
