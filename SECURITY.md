# Security Policy

Obrigado por ajudar a manter o s-go seguro. Esta política explica como reportar vulnerabilidades, nosso processo de resposta e o que está ou não em escopo.

Nota: a versão em inglês está no final.

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
LTS tem suporte por **6 meses após o release**.
Não‑LTS têm suporte por **50 dias após o release**.
Exemplos: release em **2026-02-06** → LTS até **2026-08-06**; não‑LTS até **2026-03-28**.

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x (LTS) | :white_check_mark: |
| 2.7.x (inclui 2.7.2 bugfix) | :white_check_mark: |
| 1.0.x       | :white_check_mark: |
| 0.0.0       | :white_check_mark: |

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

---

# Security Policy (English)

Thanks for helping keep s-go secure. This policy explains how to report vulnerabilities, our response process, and what is in or out of scope.

## How to report
Please **do not** open public issues for vulnerabilities.

Use the GitHub Security Advisory:
https://github.com/devcombiscoito/s-go/security/advisories/new

If you cannot access the advisory, open a **non-sensitive** issue and request private contact.

## What to include
Provide as much as possible:
- s-go version (`s version`)
- OS and version
- Architecture (e.g., amd64, arm64)
- Detailed steps to reproduce
- Expected vs actual results
- Potential impact (e.g., RCE, data leak)
- Relevant logs/output (if any)
- POC (if available), with safe instructions

## Response SLA
- **Acknowledgement**: within 72 hours
- **Status updates**: weekly while actively investigating
- **Fix**: as fast as possible, based on severity

## Disclosure process
We follow responsible disclosure:
1. Receive and assess severity
2. Reproduce and confirm
3. Develop and test a fix
4. Release the fix and communicate the vulnerability

## Severity
General criteria:
- **Critical**: remote code execution, broad data exposure
- **High**: local security bypass, privilege escalation
- **Medium**: limited exposure, meaningful denial of service
- **Low**: configuration issues with limited impact

## Supported versions
LTS is supported for **6 months after release**.
Non‑LTS is supported for **50 days after release**.
Examples: released on **2026-02-06** → LTS until **2026-08-06**; non‑LTS until **2026-03-28**.

| Version | Supported          |
| ------- | ------------------ |
| 2.0.x (LTS) | :white_check_mark: |
| 2.7.x (includes 2.7.2 bugfix) | :white_check_mark: |
| 1.0.x       | :white_check_mark: |
| 0.0.0       | :white_check_mark: |

## Scope
Includes:
- `s-go` CLI
- Scripts and templates in this repository
- Automatic updates of the CLI itself

Excludes:
- Third-party dependencies
- User local environment
- Issues in external tools or operating systems

## Reporting best practices
- Do not exploit beyond what is needed to prove impact
- Avoid accessing data that is not yours
- Do not publish POCs before a fix is released

## Thanks
Security contributions are welcome and may be acknowledged in release notes.
