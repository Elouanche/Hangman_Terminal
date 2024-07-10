package main

import (
	"bufio"
	"os"
)

//Ouvre les fichiers où les mots sont enregistré
func lireMotsDuFichier(dossierMots string) ([]string, error) {
	fichier, err := os.Open(dossierMots)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()

	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mots, nil
}
