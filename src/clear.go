package main

import (
	"os"
	"os/exec"
	"runtime"
)

//permets de supprimer les éléments affichés dans le terminal
func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
