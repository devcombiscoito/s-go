package app

import (
	"fmt"
	"os"
	"s-go/internal/commands"
	"s-go/internal/commands/help"
	"s-go/internal/commands/runners"
)

func Verify() {
	CheckUpdate()

	if len(os.Args) < 2 {
		fmt.Println("Por favor, forneça um argumento.")
		return
	}

	arg := os.Args[1]

	if arg == "go" || arg == "-g" {
		runners.RodarGo()
	} else if arg == "html" || arg == "-h" {
		runners.RodarHtml()
	} else if arg == "help" || arg == "--h" {
		help.Help()
	} else if arg == "python" || arg == "-p" {
		runners.RodarPy()
	} else if arg == "update" || arg == "--u" {
		commands.RodarUpdate()
	} else if arg == "install" || arg == "-i" {
		if len(os.Args) > 2 && os.Args[2] == "-y" {
			commands.NoAskInstall()
		} else {
			commands.AskInstall()
		}
	} else if arg == "uninstall" || arg == "-u" {
		if len(os.Args) > 2 && os.Args[2] == "-y" {
			commands.NoAskUninstall()
		} else {
			commands.AskUninstall()
		}
	} else {
		fmt.Println(`Argumento não identificado. Escreva "s help" para ajuda.`)
	}
}