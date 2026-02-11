package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func NoAskUninstall() {
	fmt.Println("==> Removendo o comando 's' do sistema...")

	installPath := installBinaryPath()
	if isWindows() {
		if err := os.Remove(installPath); err != nil {
			fmt.Println("Falha ao remover o binário:", err)
			fmt.Println("Tentando remoção forçada...")
			if err2 := forceRemoveWindowsBinary(installPath); err2 != nil {
				fmt.Println("Falha na remoção forçada:", err2)
				fmt.Println("Dica: feche terminais/editores que possam estar usando 's' e tente novamente.")
				return
			}
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

// forceRemoveWindowsBinary tenta remover o binário usando PowerShell,
// útil em casos de erro de permissão simples no os.Remove.
func forceRemoveWindowsBinary(path string) error {
	// Escapa aspas simples para uso seguro no PowerShell.
	safePath := strings.ReplaceAll(path, "'", "''")
	script := fmt.Sprintf(
		"if (Test-Path -LiteralPath '%s') { Remove-Item -LiteralPath '%s' -Force }",
		safePath, safePath,
	)

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
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
