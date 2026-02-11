package commands

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func binaryName() string {
	if isWindows() {
		return "s.exe"
	}
	return "s"
}

func installDir() string {
	if isWindows() {
		base := os.Getenv("LOCALAPPDATA")
		if base == "" {
			if home, err := os.UserHomeDir(); err == nil {
				base = filepath.Join(home, "AppData", "Local")
			}
		}
		return filepath.Join(base, "s-go", "bin")
	}
	return "/usr/local/bin"
}

func installBinaryPath() string {
	return filepath.Join(installDir(), binaryName())
}

func ensureInstallDir() error {
	return os.MkdirAll(installDir(), 0755)
}

func moveFile(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	if err := out.Close(); err != nil {
		return err
	}
	return os.Remove(src)
}

// ensurePathOnWindows garante que o diretório de instalação esteja no PATH do usuário.
// Ele atualiza apenas o PATH de usuário (não o de sistema) usando PowerShell.
func ensurePathOnWindows() {
	if !isWindows() {
		return
	}

	dir := installDir()
	if dir == "" {
		return
	}

	// Script PowerShell:
	// - Lê o PATH de usuário
	// - Se estiver vazio, define apenas com o dir
	// - Se já existir, adiciona se ainda não estiver presente
	script := fmt.Sprintf(`
$p = [Environment]::GetEnvironmentVariable('Path', 'User')
if (-not $p) {
  $p = '%s'
} else {
  $parts = $p -split ';'
  if ($parts -notcontains '%s') {
    $p = $p + ';%s'
  }
}
[Environment]::SetEnvironmentVariable('Path', $p, 'User')
`, dir, dir, dir)

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err == nil {
		fmt.Println("==> PATH do usuário atualizado. Abra um novo terminal para usar 's'.")
	}
}
