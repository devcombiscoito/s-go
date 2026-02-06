package runners

import (
	"fmt"
	"os"
	"path/filepath"
)

func RodarPy() {
	RodarPyAt("")
}

func RodarPyAt(baseDir string) {
	var projectName string

	fmt.Print("Qual o nome do projeto Python? ")
	fmt.Scan(&projectName)

	projectPath := projectName
	if baseDir != "" {
		projectPath = filepath.Join(baseDir, projectName)
	}

	os.Mkdir(projectPath, 0755)

	main := `print("Hello, World")`

	os.WriteFile(filepath.Join(projectPath, "main.py"), []byte(main), 0644)

	fmt.Println("🚀 Tudo pronto!")
}
