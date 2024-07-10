package main

import (
	"math/rand"
	"time"
)

//Charge un mot aléatoirement
func choisirMotAleatoire(mots []string) string {
	rand.Seed(time.Now().UnixNano())
	indiceAleatoire := rand.Intn(len(mots))
	return mots[indiceAleatoire]
}
