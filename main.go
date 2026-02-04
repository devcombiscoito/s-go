package main

import (
	"os"
	"os/exec"
)
func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	verify()
}