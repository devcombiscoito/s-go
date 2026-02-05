package runners

import (
	"fmt"
	"os"
	"os/exec"
)

func RodarGo() {
	var projectName string
	
	fmt.Print("Qual o nome do projeto Go? ")
	fmt.Scan(&projectName)

	os.Mkdir(projectName, 0755)

	execCommand := exec.Command("go", "mod", "init", projectName)
	execCommand.Dir = projectName
	execCommand.Run()
	
	main := `package main
	
import (
	"fmt"
)
		
func main() {
	fmt.Print("Hello, world!")
}`

	os.WriteFile(projectName+"/main.go",  []byte(main), 0644)

	fmt.Println("🚀 Tudo pronto! Arquivo main.go criado.")
}