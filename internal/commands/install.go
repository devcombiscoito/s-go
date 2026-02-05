package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func NoAskInstall() {
	fmt.Println("==> Compilando e instalando...")
	
	build := exec.Command("go", "build", "-o", "s", "./cmd/s")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	build.Run()

	os.Chmod("s", 0755)

	move := exec.Command("sudo", "mv", "s", "/usr/local/bin/s")
	move.Stdout = os.Stdout
	move.Stderr = os.Stderr
	move.Stdin = os.Stdin
	move.Run()

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
