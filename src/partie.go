package main

import "fmt"

//Fonction principale du jeu
func jouerPartie(mots []string) {
	motAleatoire := choisirMotAleatoire(mots)
	_, motMasque := afficherLettreAleatoire(motAleatoire)
	tailleMot := len(motAleatoire)
	tentatives := 10

	for tentatives > 0 {
		if motMasque == motAleatoire {
			clearTerminal()
			fmt.Println("Bravo tu as trouvé le mot ! C'était bien :", motAleatoire)
			break
		} else {
			afficherEtat(motMasque, tentatives)
			var choix string
			fmt.Scan(&choix)

			if len(choix) == tailleMot {
				if traiterMotDevine(motAleatoire, choix, &tentatives) {
					break
				}
			} else if len(choix) == 1 {
				motMasque = traiterLettre(motAleatoire, motMasque, choix, &tentatives)
			} else {
				clearTerminal()
				fmt.Println("Choix invalide.")
			}
		}

		if tentatives == 0 {
			clearTerminal()
			bonhomme(&tentatives)
			fmt.Println("Vous avez épuisé toutes vos tentatives. Le mot était:", motAleatoire)
		}
	}
}
