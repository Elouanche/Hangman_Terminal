package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//Permets d'afficher l'indice du début
func afficherLettreAleatoire(mot string) (rune, string) {
	rand.Seed(time.Now().UnixNano())
	indiceAleatoire := rand.Intn(len(mot))

	lettreAleatoire := rune(mot[indiceAleatoire])

	// Créer un mot masqué avec des underscores
	motMasque := strings.Map(func(r rune) rune {
		if r == lettreAleatoire {
			return r
		}
		return '_'
	}, mot)

	fmt.Println("Indice:", motMasque)

	return lettreAleatoire, motMasque
}
