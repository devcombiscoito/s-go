package runners

import (
	"fmt"
	"os"

)

func RodarHtml() {
	
	var projectName string
		
	fmt.Print("Qual o nome do projeto Html? ")
	fmt.Scan(&projectName)

	os.Mkdir(projectName, 0755)

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="styles.css">
    <title>	 </title>
</head>
<body>
    <h1>Hello, World!</h1>
	<script src="script.js"></script>
</body>
</html>`

	css :=  `* {
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

	os.WriteFile(projectName+"/index.html",  []byte(html), 0644)
	os.WriteFile(projectName+"/styles.css",  []byte(css), 0644)
	os.WriteFile(projectName+"/script.js", []byte(js), 0644)
	fmt.Println("🚀 Tudo pronto!")
}