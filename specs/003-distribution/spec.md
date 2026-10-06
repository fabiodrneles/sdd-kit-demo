# 003 — Distribuição

- **Prioridade:** P2
- **Status:** Done — entregue na `v1.0.0`
- **Código afetado:** `.goreleaser.yml`, `.github/workflows/`, `action.yml`

## Requisitos funcionais

- **FR-1 (D4)** Tags `v*` MUST publicar binários para Linux, macOS e Windows com checksums, depois de `make ci` verde.
- **FR-2 (D4)** A próxima versão MUST ser calculada pelo go-release-manager a partir dos commits.
- **FR-3** `go install github.com/fabiodrneles/sdd-kit-demo@latest` MUST instalar o `linkcheck`.
- **FR-4** O repositório SHOULD oferecer uma GitHub Action que roda o `linkcheck` num repositório.

## Critérios de aceite

- **AC-1** Dado o `.goreleaser.yml`, quando o CI roda `goreleaser check`, então passa.
- **AC-2** Dado um workflow de teste que usa a Action local num diretório com um link quebrado, quando roda, então o job falha.

## Decisões

- D4 respondida pelo dono em 2026-10-02 conforme as recomendações de [ANALYSIS.md §7](../ANALYSIS.md#7-decisões-respondidas-pelo-dono-em-2026-10-02).
