# Análise e decisões iniciais do sdd-kit-demo

Repositório novo, criado em 2026-10-02 para demonstrar o [sdd-kit](https://github.com/fabiodrneles/sdd-kit) de ponta a ponta: da primeira spec à primeira release, com cada passo registrado em issues e PRs.

## 1. Resumo executivo

O projeto é o `linkcheck`, uma CLI em Go que encontra **links quebrados em arquivos Markdown**: links relativos para arquivos que não existem, âncoras (`#secao`) que não existem no arquivo de destino e, opcionalmente, URLs externas que não respondem. É pequena, útil em qualquer repositório com documentação e fácil de verificar em CI.

## 2. O que foi verificado

| Comando | Resultado |
|---|---|
| `git log` | só o commit inicial com o `README.md` |
| `adopt.sh --lang go` (sdd-kit `v0.1.0`) | 24 arquivos criados |

## 3. Observações por severidade

Não há código; as observações são do próprio kit, registradas em fabiodrneles/sdd-kit#51:

### Médias

- **M1** O template Go não traz esqueleto (`go.mod`, `main.go`): o `make ci` de um repositório vazio não tem o que compilar. → resolvido neste PR com um esqueleto mínimo.

## 4. Pontos positivos (manter)

- Nada a preservar ainda.

## 5. Avaliação do README

Só o título. Será escrito na Fase 1, com comandos verificados no CI.

## 6. Melhorias recomendadas (priorizadas)

1. **Fase 1 (P0) → `v0.1.0`:** links locais e âncoras, saída `arquivo:linha`, código de saída.
2. **Fase 2 (P1) → `v0.2.0`:** URLs externas (opcional), saída JSON, configuração de exclusões.
3. **Fase 3 (P2) → `v1.0.0`:** binários por GoReleaser, versão calculada pelo go-release-manager, GitHub Action.

## 7. Decisões em aberto

| ID | Pergunta | Opções | Recomendação |
|---|---|---|---|
| D1 | URLs externas (`http(s)://`) são verificadas por padrão? | (a) não: só com `--external`; (b) sim, sempre | **(a)**: a verificação padrão fica determinística e roda offline no CI |
| D2 | Âncoras (`arquivo.md#secao`) são verificadas? | (a) sim, com os slugs de título do GitHub; (b) não | **(a)**: âncora quebrada é o link quebrado mais comum em documentação |
| D3 | Formato da saída? | (a) `arquivo:linha: mensagem` (como compiladores, clicável nos editores) e `--format json`; (b) só texto livre | **(a)** |
| D4 | Como versionar e publicar? | (a) GoReleaser com a versão calculada pelo go-release-manager (Conventional Commits); (b) tags manuais | **(a)**: demonstra o plugin `sdd-release` do kit |
