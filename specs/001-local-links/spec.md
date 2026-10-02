# 001 — Links locais e âncoras

- **Prioridade:** P0
- **Status:** Approved — decisões respondidas pelo dono em 2026-10-02
- **Código afetado:** `main.go`, `internal/`

## Contexto

O caso principal: documentação com links relativos para arquivos que mudaram de nome ou para seções que não existem mais.

## Requisitos funcionais

- **FR-1** `linkcheck [CAMINHO...]` MUST verificar todos os arquivos `.md` dos caminhos (diretórios recursivamente; padrão `.`), ignorando `.git/` e `node_modules/`.
- **FR-2** MUST reconhecer links inline (`[texto](destino)`), imagens (`![alt](destino)`) e definições de referência (`[id]: destino`), fora de blocos de código.
- **FR-3** Um destino relativo MUST existir no disco, resolvido a partir do diretório do arquivo.
- **FR-4 (D2)** Uma âncora (`#secao`, local ou `arquivo.md#secao`) MUST corresponder a um título do arquivo de destino, com o slug do GitHub (minúsculas, espaços viram `-`, pontuação removida, títulos repetidos ganham `-1`, `-2`).
- **FR-5 (D3)** Cada problema MUST ser impresso como `arquivo:linha: motivo: destino`, ordenado por arquivo e linha.
- **FR-6** O código de saída MUST ser 0 sem problemas, 1 com links quebrados e 2 em erro de uso.

## Critérios de aceite

- **AC-1** Dado `a.md` com `[x](b.md)` e sem `b.md`, quando `linkcheck` roda, então imprime `a.md:1: arquivo não encontrado: b.md` e sai com 1.
- **AC-2** Dado `a.md` com `[x](#nao-existe)`, quando `linkcheck` roda, então aponta a âncora e sai com 1.
- **AC-3** Dado `a.md` com `[x](b.md#instalação)` e `b.md` com `## Instalação`, quando `linkcheck` roda, então sai com 0.
- **AC-4** Dado um link quebrado dentro de um bloco de código cercado, quando `linkcheck` roda, então ele é ignorado.
- **AC-5** Dado um diretório só com links válidos, quando `linkcheck` roda, então não imprime nada e sai com 0.
- **AC-6** Dada uma flag desconhecida, quando `linkcheck` roda, então sai com 2.

## Fora de escopo

- URLs externas (spec 002).

## Mudanças

- MODIFIED AC-3: a âncora de `## Instalação` é `#instalação`. O slug do GitHub mantém os acentos; o texto anterior, `#instalacao`, contradizia o FR-4 (T2, #6).

## Decisões

- D2 e D3 respondidas pelo dono em 2026-10-02 conforme as recomendações de [ANALYSIS.md §7](../ANALYSIS.md#7-decisões-respondidas-pelo-dono-em-2026-10-02).
