package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RodarUpdate() {
	directory := "/usr/local/bin"

	fmt.Println("==> Buscando atualizações no GitHub...")
	execPull := exec.Command("git", "pull", "main", "main")
	execPull.Stdout = os.Stdout
	execPull.Stderr = os.Stderr
	execPull.Run()

	fmt.Println("==> Compilando nova versão...")
	execUpdate := exec.Command("go", "build", "-o", "s", "./cmd/s")
	execUpdate.Run()

	execMove := exec.Command("sudo", "mv", "s", directory)
	execMove.Stdin = os.Stdin
	execMove.Stdout = os.Stdout
	execMove.Stderr = os.Stderr
	execMove.Run()

	println(`Comando "s" atualizado com sucesso!`)
}
