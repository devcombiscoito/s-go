package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	DefaultPath string `json:"default_path"`
}

func configFilePath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "s-go", "config.json"), nil
}

func loadConfig() (Config, error) {
	var cfg Config
	path, err := configFilePath()
	if err != nil {
		return cfg, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return cfg, err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func saveConfig(cfg Config) error {
	path, err := configFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func ensureConfigFile() (string, error) {
	path, err := configFilePath()
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}
	if !os.IsNotExist(err) {
		return "", err
	}
	if err := saveConfig(Config{}); err != nil {
		return "", err
	}
	return path, nil
}

func ConfigList() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Erro ao carregar configuração:", err)
		return
	}
	if cfg.DefaultPath == "" {
		fmt.Println("default_path: (vazio)")
		return
	}
	fmt.Println("default_path:", cfg.DefaultPath)
}

func ConfigEdit() {
	path, err := ensureConfigFile()
	if err != nil {
		fmt.Println("Erro ao preparar configuração:", err)
		return
	}
	editor := pickEditor()
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao abrir editor:", err)
	}
}

func pickEditor() string {
	if v := os.Getenv("EDITOR"); v != "" {
		return v
	}
	if isWindows() {
		return "notepad"
	}
	return "nano"
}

func DefaultProjectDir() string {
	cfg, err := loadConfig()
	if err != nil || cfg.DefaultPath == "" {
		return ""
	}
	return cfg.DefaultPath
}
