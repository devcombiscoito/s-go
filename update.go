package main

import (
	"os/exec"
	"os"
)

func rodarUpdate() {
	directory := "/usr/local/bin"
	
	execUpdate := exec.Command("go", "build", "-o", "s")
	execUpdate.Run()


	execMove := exec.Command("sudo", "mv", "s", directory)
	execMove.Stdin = os.Stdin
	execMove.Stdout = os.Stdout
	execMove.Stderr = os.Stderr
	execMove.Run()

	println(`Comando "s" atualizado com sucesso!`)
}