package runners

import (
	"fmt"
	"os"
)

func RodarPy() {
	
	var projectName string
	
	fmt.Print("Qual o nome do projeto Python? ")
	fmt.Scan(&projectName)

	os.Mkdir(projectName, 0755)
	
	main := `print("Hello, World")`

	os.WriteFile(projectName+"/main.py",  []byte(main), 0644)

	fmt.Println("🚀 Tudo pronto!")
}