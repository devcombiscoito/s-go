package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RodarUpdate() {
	installPath := installBinaryPath()

	fmt.Println("==> Buscando atualizações no GitHub...")
	execPull := exec.Command("git", "pull", "main", "master")
	execPull.Stdout = os.Stdout
	execPull.Stderr = os.Stderr
	execPull.Run()

	fmt.Println("==> Compilando nova versão...")
	binName := binaryName()
	execUpdate := exec.Command("go", "build", "-o", binName, "./cmd/s")
	execUpdate.Run()

	if isWindows() {
		if err := ensureInstallDir(); err != nil {
			fmt.Println("Falha ao criar diretório de instalação:", err)
			return
		}
		if err := moveFile(binName, installPath); err != nil {
			fmt.Println("Falha ao mover o binário:", err)
			return
		}
	} else {
		execMove := exec.Command("sudo", "mv", binName, installPath)
		execMove.Stdin = os.Stdin
		execMove.Stdout = os.Stdout
		execMove.Stderr = os.Stderr
		execMove.Run()
	}

	println(`Comando "s" atualizado com sucesso!`)
}
