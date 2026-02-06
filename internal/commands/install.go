package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func NoAskInstall() {
	fmt.Println("==> Compilando e instalando...")

	if !ensureDependencies() {
		return
	}

	binName := binaryName()
	build := exec.Command("go", "build", "-o", binName, "./cmd/s")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	build.Run()

	if !isWindows() {
		os.Chmod(binName, 0755)
	}

	installPath := installBinaryPath()
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
		move := exec.Command("sudo", "mv", binName, installPath)
		move.Stdout = os.Stdout
		move.Stderr = os.Stderr
		move.Stdin = os.Stdin
		move.Run()
	}

	fmt.Println("==> Pronto! O comando 's' agora é global.")
}

func AskInstall() {
	var answer string
	fmt.Print("Deseja instalar o programa no sistema? [S/n]: ")
	fmt.Scanln(&answer)

	if answer == "s" || answer == "S" || answer == "sim" || answer == "Sim" || answer == "SIM" || answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" || answer == "YES" || answer == "" {
		NoAskInstall()
	} else {
		fmt.Println("Instalação cancelada.")
	}
}
