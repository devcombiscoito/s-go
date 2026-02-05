package app

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func CheckUpdate() {
	// URL do arquivo raw no GitHub
	url := "https://raw.githubusercontent.com/devcombiscoito/s-go/main/internal/app/version.go"

	client := http.Client{
		Timeout: 2 * time.Second, // Timeout curto para não travar o cli
	}

	resp, err := client.Get(url)
	if err != nil {
		// Falha silenciosa se não tiver internet ou github fora
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Procura pela string de versão usando regex
	// Espera algo como: const CurrentVersion = "1.0.1"
	re := regexp.MustCompile(`const CurrentVersion = "([^"]+)"`)
	matches := re.FindStringSubmatch(string(body))

	if len(matches) > 1 {
		remoteVersion := matches[1]
		if isNewer(remoteVersion, CurrentVersion) {
			fmt.Printf("\n🔔 \033[33mA nova versão %s está disponível! (Atual: %s)\033[0m\n", remoteVersion, CurrentVersion)
			fmt.Println("👉 Execute 's update' para atualizar.")
			fmt.Println("")
		}
	}
}

// isNewer verifica se v1 é maior que v2.
func isNewer(remote, current string) bool {
	// Remove possível prefixo "v"
	remote = strings.TrimPrefix(remote, "v")
	current = strings.TrimPrefix(current, "v")

	v1Parts := strings.Split(remote, ".")
	v2Parts := strings.Split(current, ".")

	// Normaliza para ter o mesmo número de partes (até 3)
	for len(v1Parts) < 3 {
		v1Parts = append(v1Parts, "0")
	}
	for len(v2Parts) < 3 {
		v2Parts = append(v2Parts, "0")
	}

	for i := 0; i < 3; i++ {
		var n1, n2 int
		fmt.Sscanf(v1Parts[i], "%d", &n1)
		fmt.Sscanf(v2Parts[i], "%d", &n2)

		if n1 > n2 {
			return true
		}
		if n1 < n2 {
			return false
		}
	}
	return false
}
