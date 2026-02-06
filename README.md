# s-go 🚀

O `s-go` é uma ferramenta CLI (Command Line Interface) projetada para automatizar a criação de arquivos boilerplate e estruturas de projetos para Go, Python e Web (HTML/CSS/JS), ajudando você a economizar tempo em configurações repetitivas.

## 🛠️ Instalação

Siga os passos abaixo para compilar e instalar o `s-go` no seu sistema.
O comando `install` agora verifica dependências (`go`, `git`, `python`) e oferece instalação automática.

1. **Clone o repositório** (caso ainda não tenha feito):
   ```bash
   git clone https://github.com/devcombiscoito/s-go.git
   cd s-go
   ```

2. **Compile o projeto**:
   ```bash
   go build -o s ./cmd/s
   ```

3. **Instale no sistema**:
   - Linux/macOS: move o executável para `/usr/local/bin` (pode pedir senha de sudo).
   - Windows: move o executável para `%LOCALAPPDATA%\\s-go\\bin` (garanta esse caminho no `PATH`).
   ```bash
   ./s install
   ```

### Instalação via Release
Você também pode baixar os binários prontos na página de Releases do GitHub:
- https://github.com/devcombiscoito/s-go/releases
- `s` (Linux)
- `s.exe` (Windows)
- `s-darwin-amd64` e `s-darwin-arm64` (macOS)

### Checksums
Para gerar os checksums dos binários:
```bash
./scripts/checksums.sh s s.exe s-darwin-amd64 s-darwin-arm64
```
Isso cria o arquivo `SHA256SUMS`.

## 💻 Comandos Disponíveis

Aqui estão os comandos que você pode usar para agilizar seu desenvolvimento:

| Comando | Atalho | Descrição |
|---------|--------|-----------|
| `s go` | `s -g` | Cria um novo projeto Go, inicializa o `go.mod` e cria um `main.go`. |
| `s python` | `s -p` | Gera um script Python básico (`main.py`) pronto para rodar. |
| `s html` | `s -h` | Cria uma estrutura Web completa com `index.html`, `styles.css` e `script.js`. |
| `s update` | `s --u` | Atualiza a ferramenta compilando a versão mais recente do código fonte. |
| `s version` | `s --v` | Exibe a versão instalada do `s-go`. |
| `s help` | `s --h` | Exibe a lista de comandos e ajuda. |
| `s uninstall` | `s -u`| Remove o comando `s` do seu sistema. |

### Dicas Úteis
- **Instalação/Desinstalação Rápida**: Adicione a flag `-y` aos comandos `install` ou `uninstall` para pular a confirmação (ex: `./s install -y`).
- **Dependências**: Se `go`, `git` ou `python` não estiverem instalados, o `install` oferece instalação automática no Linux, Windows e macOS.

## 📁 Estrutura do Projeto

- **cmd/s**: Ponto de entrada da aplicação (main).
- **internal/app**: Lógica de verificação e roteamento de comandos.
- **internal/commands**: Implementação das funcionalidades de instalação, atualização e help.
- **internal/commands/runners**: Lógica específica para gerar os arquivos de cada linguagem (Go, Python, HTML).

---
Feito com dedicação por Lorenzo. Aproveite!
