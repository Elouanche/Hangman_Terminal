package main

import (
	"fmt"
	"strings"
)

//Verifications de la lettre entré
func traiterLettre(motAleatoire, motMasque, choix string, tentatives *int) string {
	if !isLettre(choix) {
		clearTerminal()
		fmt.Println("Veuillez choisir une lettre.")
		return motMasque
	}

	lettreJoueur := rune(choix[0])

	if lettresDejaDevinees[lettreJoueur] {
		clearTerminal()
		fmt.Printf("La lettre '%c' a déjà été devinée.\n", lettreJoueur)
	} else if lettresDevinéesIncorrectes[lettreJoueur] {
		clearTerminal()
		fmt.Printf("La lettre '%c' a déjà été essayée et n'est pas dans le mot.\n", lettreJoueur)
	} else if strings.ContainsRune(motMasque, lettreJoueur) {
		clearTerminal()
		fmt.Printf("La lettre '%c' fait partie du mot masqué initial et ne peut pas être devinée.\n", lettreJoueur)
	} else if strings.ContainsRune(motAleatoire, lettreJoueur) {
		clearTerminal()
		fmt.Printf("La lettre '%c' est dans le mot!\n", lettreJoueur)
		lettresDejaDevinees[lettreJoueur] = true
		motMasque = strings.Map(func(r rune) rune {
			if r == lettreJoueur || r != '_' && strings.ContainsRune(motMasque, r) {
				return r
			}
			return '_'
		}, motAleatoire)
	} else {
		clearTerminal()
		fmt.Printf("La lettre '%c' n'est pas dans le mot.\n", lettreJoueur)
		(*tentatives)--
		bonhomme(tentatives)
		lettresDevinéesIncorrectes[lettreJoueur] = true
	}

	return motMasque
}
