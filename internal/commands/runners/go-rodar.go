package runners

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RodarGo() {
	RodarGoAt("")
}

func RodarGoAt(baseDir string) {
	var projectName string

	fmt.Print("Qual o nome do projeto Go? ")
	fmt.Scan(&projectName)

	projectPath := projectName
	if baseDir != "" {
		projectPath = filepath.Join(baseDir, projectName)
	}

	os.Mkdir(projectPath, 0755)

	execCommand := exec.Command("go", "mod", "init", projectName)
	execCommand.Dir = projectPath
	execCommand.Run()

	main := `package main
	
import (
	"fmt"
)
		
func main() {
	fmt.Print("Hello, world!")
}`

	os.WriteFile(filepath.Join(projectPath, "main.go"), []byte(main), 0644)

	fmt.Println("🚀 Tudo pronto! Arquivo main.go criado.")
}
