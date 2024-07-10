package main

import (
	"bufio"
	"fmt"
	"os"
)

//permets de lire le fichier hangman.txt
func lireFichierHangman(startLine int, endLine int) ([]string, error) {
	fichier, err := os.Open("./ressources/hangman.txt")
	if err != nil {
		return nil, err
	}
	defer fichier.Close()

	var positions []string
	scanner := bufio.NewScanner(fichier)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		if lineNumber >= startLine && lineNumber <= endLine {
			positions = append(positions, scanner.Text())
		}
		if lineNumber > endLine {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if startLine > lineNumber || endLine > lineNumber {
		return nil, fmt.Errorf("PARTIE TERMINE")
	}
	return positions, nil
}

//Permets d'afficher le pendu
func bonhomme(tentatives *int) {
	linesRange := map[int][2]int{
		9: {1, 8},
		8: {9, 16},
		7: {17, 24},
		6: {25, 32},
		5: {33, 40},
		4: {41, 48},
		3: {49, 56},
		2: {57, 64},
		1: {65, 72},
		0: {73, 80},
	}

	if rangeValues, ok := linesRange[*tentatives]; ok {
		startLine, endLine := rangeValues[0], rangeValues[1]
		lines, err := lireFichierHangman(startLine, endLine)
		if err != nil {
			fmt.Println(err) 
			return
		}
		for _, lignes := range lines {
			fmt.Println(lignes)
		}
	} else {
		fmt.Println("Nombre de tentatives invalide.")
	}
}
