# Changelog
Todas as mudanças relevantes deste projeto serão documentadas aqui.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/)
e este projeto segue [SemVer](https://semver.org/lang/pt-BR/).

Nota: a versão em inglês está no final.

## [2.0.0] - 2026-02-06 - Grande Atualização!!! (LTS)
### Adicionado
- Verificação e instalação automática de dependências (`go`, `git`, `python`) durante `install`.
- Suporte de instalação em Windows e macOS.
- Binários de release para Linux, Windows e macOS.

### Alterado
- `install`, `update` e `uninstall` agora usam caminhos específicos por sistema operacional.
- Documentação de instalação e releases atualizada.

### Removido
- Arquivos e pastas antigas de testes/rascunhos.

## [1.0.0] - 2026-02-05
### Adicionado
- CLI inicial para gerar projetos Go, Python e HTML.
- Comandos `install`, `uninstall`, `update`, `version` e `help`.

[2.0.0]: https://github.com/devcombiscoito/s-go/releases/tag/v2.0.0

---

# Changelog (English)
All notable changes to this project will be documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [SemVer](https://semver.org/).

## [2.0.0] - 2026-02-06 - Major Update!!! (LTS)
### Added
- Dependency check and automatic install (`go`, `git`, `python`) during `install`.
- Installation support for Windows and macOS.
- Release binaries for Linux, Windows, and macOS.

### Changed
- `install`, `update`, and `uninstall` now use OS-specific paths.
- Installation and release docs updated.

### Removed
- Old test/draft folders and files.

## [1.0.0] - 2026-02-05
### Added
- Initial CLI to generate Go, Python, and HTML projects.
- `install`, `uninstall`, `update`, `version`, and `help` commands.

[2.0.0]: https://github.com/devcombiscoito/s-go/releases/tag/v2.0.0
