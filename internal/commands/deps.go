package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type dependency struct {
	name string
}

func ensureDependencies() bool {
	missing := missingDependencies()
	if len(missing) == 0 {
		return true
	}

	fmt.Println("Dependências ausentes:")
	for i, dep := range missing {
		fmt.Printf("  %d. %s\n", i+1, dep.name)
	}

	var answer string
	fmt.Print("Deseja instalar agora? [S/n]: ")
	fmt.Scanln(&answer)
	if isNo(answer) {
		fmt.Println("Instalação cancelada.")
		return false
	}

	selected := selectDependencies(missing)
	if len(selected) == 0 {
		fmt.Println("Nenhuma dependência selecionada. Cancelando.")
		return false
	}

	for _, dep := range selected {
		if err := installDependency(dep); err != nil {
			fmt.Printf("Falha ao instalar %s: %v\n", dep.name, err)
			return false
		}
	}

	return true
}

func missingDependencies() []dependency {
	var missing []dependency

	if _, err := exec.LookPath("go"); err != nil {
		missing = append(missing, dependency{name: "go"})
	}
	if _, err := exec.LookPath("git"); err != nil {
		missing = append(missing, dependency{name: "git"})
	}
	if !hasPython() {
		missing = append(missing, dependency{name: "python"})
	}

	return missing
}

func hasPython() bool {
	if _, err := exec.LookPath("python"); err == nil {
		return true
	}
	if _, err := exec.LookPath("python3"); err == nil {
		return true
	}
	if runtime.GOOS == "windows" {
		if _, err := exec.LookPath("py"); err == nil {
			return true
		}
	}
	return false
}

func selectDependencies(missing []dependency) []dependency {
	if len(missing) == 1 {
		return missing
	}

	var input string
	fmt.Print("Digite os números separados por vírgula (ENTER para todas): ")
	fmt.Scanln(&input)

	input = strings.TrimSpace(input)
	if input == "" {
		return missing
	}

	parts := strings.Split(input, ",")
	seen := map[int]bool{}
	var selected []dependency
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		n, err := strconv.Atoi(part)
		if err != nil || n < 1 || n > len(missing) || seen[n] {
			continue
		}
		seen[n] = true
		selected = append(selected, missing[n-1])
	}

	return selected
}

func installDependency(dep dependency) error {
	switch runtime.GOOS {
	case "windows":
		return installWindows(dep.name)
	case "darwin":
		return installMac(dep.name)
	default:
		return installLinux(dep.name)
	}
}

func installLinux(name string) error {
	manager := detectLinuxPackageManager()
	if manager == "" {
		return fmt.Errorf("nenhum gerenciador de pacotes encontrado (apt, dnf, yum, pacman, apk, zypper)")
	}

	pkg := linuxPackageName(manager, name)
	if pkg == "" {
		return fmt.Errorf("pacote não mapeado para %s", name)
	}

	var cmd []string
	switch manager {
	case "apt-get":
		cmd = []string{"apt-get", "install", "-y", pkg}
	case "dnf":
		cmd = []string{"dnf", "install", "-y", pkg}
	case "yum":
		cmd = []string{"yum", "install", "-y", pkg}
	case "pacman":
		cmd = []string{"pacman", "-S", "--noconfirm", pkg}
	case "apk":
		cmd = []string{"apk", "add", pkg}
	case "zypper":
		cmd = []string{"zypper", "--non-interactive", "install", pkg}
	default:
		return fmt.Errorf("gerenciador não suportado: %s", manager)
	}

	return runWithOptionalSudo(cmd)
}

func detectLinuxPackageManager() string {
	candidates := []string{"apt-get", "dnf", "yum", "pacman", "apk", "zypper"}
	for _, c := range candidates {
		if _, err := exec.LookPath(c); err == nil {
			return c
		}
	}
	return ""
}

func linuxPackageName(manager, name string) string {
	switch manager {
	case "apt-get":
		switch name {
		case "go":
			return "golang-go"
		case "git":
			return "git"
		case "python":
			return "python3"
		}
	default:
		switch name {
		case "go":
			return "go"
		case "git":
			return "git"
		case "python":
			return "python3"
		}
	}
	return ""
}

func installMac(name string) error {
	if _, err := exec.LookPath("brew"); err != nil {
		return fmt.Errorf("Homebrew não encontrado. Instale o brew e tente novamente")
	}

	var pkg string
	switch name {
	case "go":
		pkg = "go"
	case "git":
		pkg = "git"
	case "python":
		pkg = "python"
	default:
		return fmt.Errorf("pacote não mapeado para %s", name)
	}

	cmd := []string{"brew", "install", pkg}
	return runCommand(cmd)
}

func installWindows(name string) error {
	if _, err := exec.LookPath("winget"); err == nil {
		id := wingetID(name)
		if id == "" {
			return fmt.Errorf("pacote não mapeado para %s", name)
		}
		cmd := []string{"winget", "install", "-e", "--id", id}
		return runCommand(cmd)
	}

	if _, err := exec.LookPath("choco"); err == nil {
		pkg := chocoPackage(name)
		if pkg == "" {
			return fmt.Errorf("pacote não mapeado para %s", name)
		}
		cmd := []string{"choco", "install", "-y", pkg}
		return runCommand(cmd)
	}

	manager, err := ensureWindowsPackageManager()
	if err != nil {
		return err
	}

	switch manager {
	case "winget":
		id := wingetID(name)
		if id == "" {
			return fmt.Errorf("pacote não mapeado para %s", name)
		}
		cmd := []string{"winget", "install", "-e", "--id", id}
		return runCommand(cmd)
	case "choco":
		pkg := chocoPackage(name)
		if pkg == "" {
			return fmt.Errorf("pacote não mapeado para %s", name)
		}
		cmd := []string{"choco", "install", "-y", pkg}
		return runCommand(cmd)
	default:
		return fmt.Errorf("nenhum gerenciador disponível")
	}
}

func wingetID(name string) string {
	switch name {
	case "go":
		return "GoLang.Go"
	case "git":
		return "Git.Git"
	case "python":
		return "Python.Python.3"
	}
	return ""
}

func chocoPackage(name string) string {
	switch name {
	case "go":
		return "golang"
	case "git":
		return "git"
	case "python":
		return "python"
	}
	return ""
}

func ensureWindowsPackageManager() (string, error) {
	fmt.Println("Nenhum gerenciador encontrado (winget ou choco).")
	fmt.Println("Selecione um para instalar:")
	fmt.Println("  1) winget (Microsoft App Installer)")
	fmt.Println("  2) choco (Chocolatey)")

	var choice string
	fmt.Print("Opcao [1/2]: ")
	fmt.Scanln(&choice)

	switch strings.TrimSpace(choice) {
	case "1":
		if err := installWinget(); err != nil {
			return "", err
		}
		if _, err := exec.LookPath("winget"); err == nil {
			return "winget", nil
		}
		return "", fmt.Errorf("winget nao encontrado apos instalacao")
	case "2":
		if err := installChoco(); err != nil {
			return "", err
		}
		if _, err := exec.LookPath("choco"); err == nil {
			return "choco", nil
		}
		return "", fmt.Errorf("choco nao encontrado apos instalacao")
	default:
		return "", fmt.Errorf("opcao invalida")
	}
}

func installWinget() error {
	cmd := []string{
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy",
		"Bypass",
		"-Command",
		"$ProgressPreference='SilentlyContinue'; " +
			"$url='https://aka.ms/getwinget'; " +
			"$out=\"$env:TEMP\\Microsoft.DesktopAppInstaller.msixbundle\"; " +
			"Invoke-WebRequest -Uri $url -OutFile $out; " +
			"Add-AppxPackage -Path $out",
	}
	return runCommand(cmd)
}

func installChoco() error {
	cmd := []string{
		"powershell",
		"-NoProfile",
		"-ExecutionPolicy",
		"Bypass",
		"-Command",
		"Set-ExecutionPolicy Bypass -Scope Process -Force; " +
			"[System.Net.ServicePointManager]::SecurityProtocol = " +
			"[System.Net.ServicePointManager]::SecurityProtocol -bor 3072; " +
			"iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))",
	}
	return runCommand(cmd)
}

func runWithOptionalSudo(cmd []string) error {
	if runtime.GOOS != "windows" {
		if _, err := exec.LookPath("sudo"); err == nil {
			cmd = append([]string{"sudo"}, cmd...)
		}
	}
	return runCommand(cmd)
}

func runCommand(cmd []string) error {
	if len(cmd) == 0 {
		return fmt.Errorf("comando vazio")
	}
	execCmd := exec.Command(cmd[0], cmd[1:]...)
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	return execCmd.Run()
}

func isNo(answer string) bool {
	switch strings.TrimSpace(strings.ToLower(answer)) {
	case "n", "nao", "não", "no":
		return true
	default:
		return false
	}
}
