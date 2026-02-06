# Contribuindo

Obrigado por querer contribuir com o s-go!

Este guia explica como propor mudanças, o fluxo recomendado e os padrões do projeto.

Nota: a versão em inglês está no final.

## Como contribuir
1. Faça um fork do projeto
2. Crie um branch a partir de `master`
3. Faça suas mudanças
4. Rode os testes: `go test ./...`
5. Abra um Pull Request

## Tipos de contribuição
- Correções de bug
- Melhorias de documentação
- Novos templates ou melhorias nos existentes
- Ajustes de UX na CLI
- Otimizações e refactors pequenos

## Configuração local
Requisitos mínimos:
- Go instalado
- Git instalado

Clonar e rodar:
```bash
git clone https://github.com/devcombiscoito/s-go.git
cd s-go
go test ./...
```

## Estilo de código
- Use `gofmt` antes de enviar PR
- Evite mudanças grandes sem abrir issue antes
- Prefira funções pequenas e bem nomeadas
- Não introduza dependências novas sem justificativa

## Testes
- Adicione testes quando fizer mudanças em lógica
- Evite testes lentos
- Garanta que `go test ./...` passa

## Commits
- Mensagens claras e curtas
- Use prefixos quando fizer sentido: `feat:`, `fix:`, `docs:`, `chore:`, `test:`

## Pull Requests
Inclua no PR:
- Resumo do que mudou
- Como testar
- Screenshots (se for algo visual)

## Reportando bugs
Use o template de bug nas issues e inclua:
- Passos para reproduzir
- Versões do Go e do s-go
- Sistema operacional

## Propondo novas features
Antes de abrir PR grande:
1. Abra uma issue descrevendo o problema
2. Sugira a solução
3. Espere feedback

---

# Contributing (English)

Thanks for contributing to s-go!

This guide explains how to propose changes, the recommended workflow, and project standards.

## How to contribute
1. Fork the repository
2. Create a branch from `master`
3. Make your changes
4. Run tests: `go test ./...`
5. Open a Pull Request

## Contribution types
- Bug fixes
- Documentation improvements
- New templates or enhancements
- CLI UX improvements
- Small optimizations and refactors

## Local setup
Minimum requirements:
- Go installed
- Git installed

Clone and run:
```bash
git clone https://github.com/devcombiscoito/s-go.git
cd s-go
go test ./...
```

## Code style
- Run `gofmt` before submitting a PR
- Avoid large changes without an issue first
- Prefer small, well-named functions
- Avoid adding dependencies without justification

## Tests
- Add tests when changing logic
- Avoid slow tests
- Ensure `go test ./...` passes

## Commits
- Use clear, short messages
- Prefixes when appropriate: `feat:`, `fix:`, `docs:`, `chore:`, `test:`

## Pull Requests
Include in the PR:
- Summary of changes
- How to test
- Screenshots (if visual)

## Reporting bugs
Use the bug template and include:
- Steps to reproduce
- Go and s-go versions
- Operating system

## Proposing features
Before opening a large PR:
1. Open an issue describing the problem
2. Suggest a solution
3. Wait for feedback
