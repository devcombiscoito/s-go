# s-go 🚀
[![CI](https://github.com/devcombiscoito/s-go/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/devcombiscoito/s-go/actions/workflows/ci.yml)

O `s-go` é uma ferramenta CLI (Command Line Interface) projetada para automatizar a criação de arquivos boilerplate e estruturas de projetos para Go, Python e Web (HTML/CSS/JS), ajudando você a economizar tempo em configurações repetitivas.

Nota: a versão em inglês está no final.

## 🛠️ Instalação

Há duas formas: **a partir do código-fonte** ou **via Release (binário pronto)**.

O comando `install` verifica dependências (`go`, `git`, `python`) e oferece instalação automática.

### Opção A — Código-fonte (recomendado para desenvolvedores)

1. **Clone o repositório**:
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

### Opção B — Release (binário pronto)
Baixe o binário na página de Releases do GitHub:
- https://github.com/devcombiscoito/s-go/releases
- `s` (Linux)
- `s.exe` (Windows)
- `s-darwin-amd64` e `s-darwin-arm64` (macOS)

Depois:
- Linux/macOS: mova o binário para `/usr/local/bin` e dê permissão:
  ```bash
  sudo mv s /usr/local/bin/s
  sudo chmod +x /usr/local/bin/s
  ```
- Windows: mova `s.exe` para `%LOCALAPPDATA%\\s-go\\bin` e adicione ao `PATH`.

Nota: o `install` foi feito para o **fluxo com código-fonte**. Se você só baixou o binário, não precisa rodar `s.exe install`.

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
| `s list` |  | Lista templates disponíveis. |
| `s init` |  | Fluxo interativo para criar projeto. |
| `s config list` |  | Lista configurações. |
| `s config edit` |  | Abre o arquivo de configuração no editor. |
| `s version` | `s --v` | Exibe a versão instalada do `s-go`. |
| `s help` | `s --h` | Exibe a lista de comandos e ajuda. |
| `s uninstall` | `s -u`| Remove o comando `s` do seu sistema. |

### Dicas Úteis
- **Instalação/Desinstalação Rápida**: Adicione a flag `-y` aos comandos `install` ou `uninstall` para pular a confirmação (ex: `./s install -y`).
- **Dependências**: Se `go`, `git` ou `python` não estiverem instalados, o `install` oferece instalação automática no Linux, Windows e macOS.

## ✅ Suporte LTS
A versão **2.7.0** é a linha **LTS** atual.
O período de suporte é de **50 dias após o release** (ex: lançado em 2026-02-06 → suporte até 2026-03-28).

## 📁 Estrutura do Projeto

- **cmd/s**: Ponto de entrada da aplicação (main).
- **internal/app**: Lógica de verificação e roteamento de comandos.
- **internal/commands**: Implementação das funcionalidades de instalação, atualização e help.
- **internal/commands/runners**: Lógica específica para gerar os arquivos de cada linguagem (Go, Python, HTML).

---
Feito com dedicação por Lorenzo. Aproveite!

---

# s-go 🚀 (English)

`s-go` is a CLI tool designed to automate boilerplate files and project structures for Go, Python, and Web (HTML/CSS/JS), helping you save time on repetitive setup.

## 🛠️ Installation

There are two ways: **from source** or **via Release (prebuilt binary)**.

The `install` command checks dependencies (`go`, `git`, `python`) and offers automatic installation.

### Option A — From source (recommended for developers)

1. **Clone the repository**:
   ```bash
   git clone https://github.com/devcombiscoito/s-go.git
   cd s-go
   ```

2. **Build the project**:
   ```bash
   go build -o s ./cmd/s
   ```

3. **Install to the system**:
   - Linux/macOS: moves the binary to `/usr/local/bin` (may require sudo).
   - Windows: moves the binary to `%LOCALAPPDATA%\\s-go\\bin` (make sure it's in your `PATH`).
   ```bash
   ./s install
   ```

### Option B — Release (prebuilt binary)
Download the binary from GitHub Releases:
- https://github.com/devcombiscoito/s-go/releases
- `s` (Linux)
- `s.exe` (Windows)
- `s-darwin-amd64` and `s-darwin-arm64` (macOS)

Then:
- Linux/macOS: move the binary to `/usr/local/bin` and make it executable:
  ```bash
  sudo mv s /usr/local/bin/s
  sudo chmod +x /usr/local/bin/s
  ```
- Windows: move `s.exe` to `%LOCALAPPDATA%\\s-go\\bin` and add it to `PATH`.

Note: `install` is meant for the **source** workflow. If you only downloaded the binary, you don't need to run `s.exe install`.

### Checksums
Generate checksums for release binaries:
```bash
./scripts/checksums.sh s s.exe s-darwin-amd64 s-darwin-arm64
```
This creates the `SHA256SUMS` file.

## 💻 Available Commands

| Command | Shortcut | Description |
|---------|----------|-------------|
| `s go` | `s -g` | Creates a new Go project, initializes `go.mod`, and creates `main.go`. |
| `s python` | `s -p` | Generates a basic Python script (`main.py`). |
| `s html` | `s -h` | Creates a web structure with `index.html`, `styles.css`, and `script.js`. |
| `s update` | `s --u` | Updates the tool by rebuilding the latest source. |
| `s list` |  | Lists available templates. |
| `s init` |  | Interactive project creation. |
| `s config list` |  | Lists configurations. |
| `s config edit` |  | Opens the config file in the editor. |
| `s version` | `s --v` | Shows the installed `s-go` version. |
| `s help` | `s --h` | Shows help and commands. |
| `s uninstall` | `s -u`| Removes the `s` command from your system. |

### Tips
- **Quick Install/Uninstall**: Add `-y` to `install` or `uninstall` to skip confirmation (ex: `./s install -y`).
- **Dependencies**: If `go`, `git`, or `python` are missing, `install` offers automatic installation on Linux, Windows, and macOS.

## ✅ LTS Support
Version **2.7.0** is the current **LTS** line.
Support lasts **50 days after release** (e.g., released on 2026-02-06 → supported until 2026-03-28).

## 📁 Project Structure

- **cmd/s**: Application entry point (main).
- **internal/app**: Command verification and routing logic.
- **internal/commands**: Install/update/help implementations.
- **internal/commands/runners**: Language-specific generators (Go, Python, HTML).

---
Made with dedication by Lorenzo. Enjoy!
