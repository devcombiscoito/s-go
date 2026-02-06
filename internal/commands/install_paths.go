package commands

import (
	"io"
	"os"
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
