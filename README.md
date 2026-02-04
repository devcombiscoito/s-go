# Projeto S: Automatizador de Workflow
Este é um utilitário desenvolvido em Go para automatizar a criação e o gerenciamento de projetos em diferentes tecnologias. Ele foi pensado para ser executado via terminal, facilitando o dia a dia de desenvolvimento.

## Funcionalidades
O comando permite interagir com diferentes stacks através de flags simples:
* Go (-g / go): Configura a estrutura base para novos projetos em Go.
* Python (-p / python): Inicializa ambientes e arquivos para Python.
* Web (-h / html): Cria a estrutura de pastas e arquivos (HTML, CSS, JS) para web design.
* Update (-u / update): O próprio comando se recompila e se atualiza no diretório de binários do sistema.
* Help (--h / help): Exibe a lista completa de comandos.

## Tecnologias Utilizadas
* Linguagem: Go (Golang)
* Ambiente: Big Linux
* Integrações: Python, HTML, CSS e JavaScript

## Como instalar
1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório.
3. No terminal, execute: go build -o s *.go
4. Mova o executável para sua pasta de binários: sudo mv s /usr/local/bin
5. Agora você pode executar o comando `s` de qualquer lugar no terminal.

## Como usar
Abra o terminal e utilize o comando `s` seguido da flag desejada. Exemplos:
* `s go` ou `s -g`: Para iniciar um projeto em Go.
* `s python` ou `s -p`: Para iniciar um projeto em Python.
* `s html` ou `s -h`: Para iniciar um projeto web.
* `s update` ou `s -u`: Para atualizar o comando S.
* `s help` ou `s --h`: Para exibir a ajuda.


