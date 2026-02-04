package main

import "fmt"

func Help() {
    fmt.Println("----------- Comando s: Seu Automatizador -----------")
    fmt.Println("Uso: s [comando]")
    fmt.Println("")
    fmt.Println("Comandos disponíveis:")
    fmt.Println("  go / -g        -> Cria um projeto Go")
    fmt.Println("  html / -h      -> Cria um projeto Web (HTML, CSS, JS)")
    fmt.Println("  python / -p    -> Cria um projeto Python")
    fmt.Println("  update / -u    -> Compila e instala a versao mais nova do s")
    fmt.Println("  help / --h     -> Mostra esta lista de comandos")
    fmt.Println("----------------------------------------------------")
}