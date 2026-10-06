# Roadmap e tarefas

Cada tarefa referencia a spec e os critérios de aceite que ela fecha. Ordem sugerida = ordem da lista.

## Fase 0 — Decisões (antes de codar)

- [x] Responder as decisões D1–D4 de [ANALYSIS.md](ANALYSIS.md) e mover as specs para `Approved`.

## Fase 1 — Funcionar de verdade (P0) → `v0.1.0`

- [x] **T1** Extração de links do Markdown (inline, imagens, referências; ignora código) — 001 FR-2, AC-4
- [x] **T2** Verificação de arquivos e âncoras, saída e códigos de saída — 001 FR-1, FR-3 a FR-6, AC-1 a AC-3, AC-5, AC-6
- [x] **T3** README com os comandos verificados no CI

## Fase 2 — Confiável (P1) → `v0.2.0`

- [x] **T4** `--external` com servidor de teste local — 002 FR-1, FR-2, AC-1, AC-2
- [x] **T5** `--format json` — 002 FR-3, AC-3
- [x] **T6** `.linkcheck.yml` com exclusões — 002 FR-4, AC-4

## Fase 3 — Profissional (P2) → `v1.0.0`

- [ ] **T7** GoReleaser e versão pelo go-release-manager — 003 FR-1 a FR-3, AC-1
- [ ] **T8** GitHub Action — 003 FR-4, AC-2
