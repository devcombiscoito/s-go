package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func NoAskUninstall() {
	fmt.Println("==> Removendo o comando 's' do sistema...")

	remove := exec.Command("sudo", "rm", "/usr/local/bin/s")
	remove.Stdout = os.Stdout
	remove.Stderr = os.Stderr
	remove.Stdin = os.Stdin
	remove.Run()

	fmt.Println("==> Desinstalado com sucesso.")
}

func AskUninstall() {
	var answer string
	fmt.Print("Tem certeza que deseja desinstalar o 's'? [s/N]: ")
	fmt.Scanln(&answer)

	if answer == "s" || answer == "S" || answer == "sim" || answer == "Sim" || answer == "SIM" || answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" || answer == "YES" || answer == "" {
		NoAskUninstall()
	} else {
		fmt.Println("Ação cancelada.")
	}
}