package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"s-go/internal/commands/runners"
)

func InitProject() {
	baseDir := ""
	if cwd, err := os.Getwd(); err == nil {
		baseDir = cwd
	}

	fmt.Println("Escolha o template:")
	fmt.Println("  1) go")
	fmt.Println("  2) python")
	fmt.Println("  3) html")
	fmt.Print("Opcao [1/2/3]: ")

	var choice string
	fmt.Scanln(&choice)
	choice = strings.TrimSpace(choice)

	if baseDir != "" {
		baseDir, _ = filepath.Abs(baseDir)
	}

	switch choice {
	case "1":
		runners.RodarGoAt(baseDir)
	case "2":
		runners.RodarPyAt(baseDir)
	case "3":
		runners.RodarHtmlAt(baseDir)
	default:
		fmt.Println("Opcao invalida")
	}
}

func ListTemplates() {
	fmt.Println("go")
	fmt.Println("python")
	fmt.Println("html")
}
