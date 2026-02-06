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

Dica: Use 's install -y' para instalar sem perguntas.
	`)
}
