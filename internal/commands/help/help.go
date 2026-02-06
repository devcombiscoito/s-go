package help

import "fmt"

func Help() {
	fmt.Println(`
Comando 's' - Seu canivete suíço de programação

Uso:
  s [comando] [opções]

Execução:
  go, -g       Roda projeto Go (run .)
  python, -p   Roda script Python
  html, -h     Abre index.html

Sistema:
  install -i    Instala o 's' globalmente (sudo no Linux/macOS)
  uninstall -u  Remove o 's' do sistema
  update, --u   Recompila e atualiza o 's'
  version, --v  Mostra a versão instalada
  help, --h     Mostra este menu

Configuração:
  list             Lista templates disponíveis
  init             Fluxo interativo para criar projeto
  config list      Lista configurações
  config edit      Abre configuração no editor

Dica: Use 's install -y' para instalar sem perguntas.
Nota: English version below.

---

Command 's' - Your programming Swiss army knife

Usage:
  s [command] [options]

Execution:
  go, -g       Run Go project (run .)
  python, -p   Run Python script
  html, -h     Open index.html

System:
  install -i    Install 's' globally (sudo on Linux/macOS)
  uninstall -u  Remove 's' from the system
  update, --u   Rebuild and update 's'
  version, --v  Show installed version
  help, --h     Show this menu

Config:
  list             List available templates
  init             Interactive project creation
  config list      List configurations
  config edit      Open config in editor

Tip: Use 's install -y' to install without prompts.
	`)
}
