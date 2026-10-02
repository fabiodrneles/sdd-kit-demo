# CLAUDE.md

Guia rápido para agentes (Claude Code) trabalharem no sdd-kit-demo sem redescobrir o projeto a cada sessão. O processo é o da skill [`sdd-delivery`](https://github.com/fabiodrneles/sdd-kit).

## Retomar o trabalho (sessão nova ou contexto perdido)

1. Leia o **comentário "Estado da fase"** mais recente no épico aberto (issues com o label `épico`): PRs, estado do CI, decisões e próximo passo.
2. Liste os **PRs abertos** e o CI de cada um, e as **issues abertas** da fase.
3. Continue do próximo passo registrado. Não refaça análise que já está em specs, issues ou PRs.

O estado do trabalho vive no GitHub, e não na conversa. Abra o ticket e o PR assim que a tarefa começar e terminar, e atualize o comentário de estado do épico a cada marco.

## O projeto

`linkcheck`: CLI em Go (só biblioteca padrão) que encontra links quebrados em arquivos Markdown. Também é o repositório de demonstração do sdd-kit, então cada passo deve ficar visível em issues e PRs.

| Caminho | O que tem |
|---|---|
| `main.go` | Ponto de entrada da CLI |
| `specs/` | Constituição, specs `NNN-nome/spec.md`, `ROADMAP.md`, `ANALYSIS.md` |

## Comandos

```text
make ci     # a mesma verificação do CI (rode antes de todo push)
make docs   # markdownlint (os links são verificados no CI)
make sdd-check  # cada AC de spec In Progress/Done citado num teste ("NNN AC-n")
```

Numa sessão na web, o hook `.claude/hooks/session-start.sh` instala as dependências e as ferramentas do CI.

## Convenções

- **Idioma:** specs, issues, PRs e documentação em português; commits e código (identificadores) em inglês.
- **Commits:** Conventional Commits (`feat:`, `fix:`, `docs:`, `test:`, `ci:`, `chore:`; `!` para mudança incompatível).
- **Branch:** uma por ticket, `<tipo>/<nº-da-issue>-<descrição>`, a partir da `main`.
- **PR:** começa com `Closes #N · Épico #M · Spec NNN` e segue o template.
- **Arquivos de status** (status das specs, checkboxes do ROADMAP, CHANGELOG) só mudam no PR de fechamento da fase.
- **Merge, tag e release** são do dono, salvo delegação explícita para uma rodada.

## Armadilhas já conhecidas

- **Toolchain baixada pelo `GOTOOLCHAIN`** não traz o `covdata`: `go test -cover` falha em pacotes sem teste. Use uma instalação completa do Go (o CI usa `setup-go`).

## Economia de uso

- Leia trechos (`sed -n 'a,bp'`, `grep -n`) em vez de arquivos inteiros, e não releia o que já leu nesta sessão.
- Para conferir CI, peça só o resumo das conclusões dos checks. Para investigar uma falha, leia o fim do log do job que falhou.
- Junte a validação num comando só (`make ci`) em vez de rodar etapas avulsas.
- Detalhes vão nos PRs e nas issues; no chat, só o resumo e o próximo passo.
