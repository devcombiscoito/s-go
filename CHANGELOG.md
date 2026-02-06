# Changelog
Todas as mudanças relevantes deste projeto serão documentadas aqui.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/)
e este projeto segue [SemVer](https://semver.org/lang/pt-BR/).

## [2.0.0] - 2026-02-06
### Adicionado
- Verificação e instalação automática de dependências (`go`, `git`, `python`) durante `install`.
- Suporte de instalação em Windows e macOS.
- Binários de release para Linux, Windows e macOS.

### Alterado
- `install`, `update` e `uninstall` agora usam caminhos específicos por sistema operacional.
- Documentação de instalação e releases atualizada.

### Removido
- Arquivos e pastas antigas de testes/rascunhos.

## [1.0.0] - 2024-??-??
### Adicionado
- CLI inicial para gerar projetos Go, Python e HTML.
- Comandos `install`, `uninstall`, `update`, `version` e `help`.

[2.0.0]: https://github.com/devcombiscoito/s-go/releases/tag/v2.0.0
