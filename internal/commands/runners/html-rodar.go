package runners

import (
	"fmt"
	"os"
	"path/filepath"
)

func RodarHtml() {
	RodarHtmlAt("")
}

func RodarHtmlAt(baseDir string) {
	var projectName string

	fmt.Print("Qual o nome do projeto Html? ")
	fmt.Scan(&projectName)

	projectPath := projectName
	if baseDir != "" {
		projectPath = filepath.Join(baseDir, projectName)
	}

	os.Mkdir(projectPath, 0755)

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="styles.css">
    <title>\t </title>
</head>
<body>
    <h1>Hello, World!</h1>
	<script src="script.js"></script>
</body>
</html>`

	css := `* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: sans-serif;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background-color: #f0f2f5;
}`

	js := `console.log("Script carregado com sucesso!");`

	os.WriteFile(filepath.Join(projectPath, "index.html"), []byte(html), 0644)
	os.WriteFile(filepath.Join(projectPath, "styles.css"), []byte(css), 0644)
	os.WriteFile(filepath.Join(projectPath, "script.js"), []byte(js), 0644)
	fmt.Println("🚀 Tudo pronto!")
}
