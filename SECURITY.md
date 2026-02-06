# Security Policy

Obrigado por ajudar a manter o s-go seguro. Esta política explica como reportar vulnerabilidades, nosso processo de resposta e o que está ou não em escopo.

## Como reportar
Por favor **não** abra issues públicas para vulnerabilidades.

Use o Security Advisory do GitHub:
https://github.com/devcombiscoito/s-go/security/advisories/new

Se não conseguir acessar o advisory, abra uma issue **sem detalhes sensíveis** e solicite contato privado.

## Informações necessárias no reporte
Inclua o máximo possível:
- Versão do s-go (`s version`)
- Sistema operacional e versão
- Arquitetura (ex: amd64, arm64)
- Passos detalhados para reproduzir
- Resultado esperado e resultado atual
- Impacto potencial (ex: execução de código, vazamento de dados)
- Logs/saídas relevantes (se houver)
- POC (se existir), com instruções seguras

## SLA de resposta
- **Confirmação de recebimento**: até 72 horas
- **Atualizações do status**: semanalmente enquanto houver investigação ativa
- **Correção**: o mais rápido possível, conforme a gravidade

## Processo de divulgação
Seguimos divulgação responsável:
1. Recebemos o relatório e avaliamos severidade
2. Reproduzimos e confirmamos a vulnerabilidade
3. Desenvolvemos e testamos a correção
4. Publicamos a correção e comunicamos a vulnerabilidade

## Severidade
Critérios gerais:
- **Crítica**: execução de código remoto, vazamento amplo de dados
- **Alta**: bypass de segurança local, escalonamento de privilégios
- **Média**: vazamento limitado, negação de serviço relevante
- **Baixa**: problemas de configuração com impacto limitado

## Versões suportadas
Mantemos suporte apenas para a versão principal mais recente.

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x   | :white_check_mark: |
| 1.0.x   | :x:                |
| 0.0.0   | :x:                |

## Escopo
Inclui:
- CLI `s-go`
- Scripts e templates distribuídos no repositório
- Atualizações automáticas do próprio CLI

Não inclui:
- Dependências externas de terceiros
- Ambiente local do usuário
- Problemas em sistemas operacionais ou ferramentas externas

## Boas práticas de reporte
- Não explore a vulnerabilidade além do necessário para provar o impacto
- Evite acessar dados que não sejam seus
- Não publique POCs antes da correção

## Agradecimentos
Contribuições de segurança são bem-vindas e serão reconhecidas nos release notes quando apropriado.
