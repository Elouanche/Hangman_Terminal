package main

import "fmt"

//Permets au joueur de choisir sa difficulté
func choisirDifficulte() {
	var x int
	fmt.Println("Quelle difficulté ? 1.Facile/2.Moyen/3.Difficile")
	fmt.Scan(&x)
	switch x {
	case 1:
		dossierMots = "./ressources/words.txt"
	case 2:
		dossierMots = "./ressources/words2.txt"
	case 3:
		dossierMots = "./ressources/words3.txt"
	default:
		fmt.Println("Difficulté non valide, choix par défaut: Facile")
		dossierMots = "./ressources/words.txt"
	}
}
