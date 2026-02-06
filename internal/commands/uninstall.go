package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func NoAskUninstall() {
	fmt.Println("==> Removendo o comando 's' do sistema...")

	installPath := installBinaryPath()
	if isWindows() {
		if err := os.Remove(installPath); err != nil {
			fmt.Println("Falha ao remover o binário:", err)
			return
		}
	} else {
		remove := exec.Command("sudo", "rm", installPath)
		remove.Stdout = os.Stdout
		remove.Stderr = os.Stderr
		remove.Stdin = os.Stdin
		remove.Run()
	}

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
