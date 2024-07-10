package main

import (
	"fmt"
	"strings"
)

//Permets aux joueurs de rejouer
func rejouer() bool {
	var choix string
	for {
		fmt.Print("Voulez-vous rejouer ? (o/n): ")
		fmt.Scan(&choix)
		if choix == "o" || choix == "n" {
			break
		} else {
			fmt.Println("Tu ne peux pas faire ça !")
		}
	}
	if strings.ToLower(choix) == "o" {
		clearTerminal()
		lettresDejaDevinees = make(map[rune]bool)
		lettresDevinéesIncorrectes = make(map[rune]bool)
		return true
	} else {
		clearTerminal()
		fmt.Println("Au revoir !")
		return false
	}
}
