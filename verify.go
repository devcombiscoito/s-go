package main

import (
    "fmt"
    "os"
)

func verify() {
    if len(os.Args) < 2 {
        fmt.Println("Por favor, forneça um argumento.")
        return
    } else if os.Args[1] == "go" || os.Args[1] == "-g"{
        rodarGo()
    } else if os.Args[1] == "html" || os.Args[1] == "-h"{
		rodarHtml()
	} else if os.Args[1] == "help" || os.Args[1] == "--h"{
		Help()
	} else if os.Args[1] == "python" || os.Args[1] == "-p"{
		rodarPy()
	} else if os.Args[1] == "update" || os.Args[1] == "-u"{
		rodarUpdate()
	} else {
		fmt.Println(`Argumento não identificado. Escreva "s help" para ajuda.`)
	}
}