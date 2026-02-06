# Contribuindo

Obrigado por querer contribuir com o s-go!

Este guia explica como propor mudanças, o fluxo recomendado e os padrões do projeto.

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
