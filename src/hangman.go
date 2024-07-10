package main

import (
	"log"
)

//Variables du jeu
var ( 
	lettresDejaDevinees = make(map[rune]bool)
 	lettresDevinéesIncorrectes = make(map[rune]bool)
 	dossierMots string
	rejouerPartie = true
)

//Initialise le jeu
func main() {
	for rejouerPartie {
		clearTerminal()
		choisirDifficulte()

		mots, err := lireMotsDuFichier(dossierMots)
		if err != nil {
			log.Fatal(err)
		}

		jouerPartie(mots)
		rejouerPartie = rejouer()
	}
}
