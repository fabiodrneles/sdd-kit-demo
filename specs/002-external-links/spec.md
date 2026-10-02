# 002 — URLs externas e saída JSON

- **Prioridade:** P1
- **Status:** Approved — decisões respondidas pelo dono em 2026-10-02
- **Código afetado:** `internal/`

## Requisitos funcionais

- **FR-1 (D1)** Com `--external`, URLs `http(s)://` MUST ser verificadas (HEAD, com GET se HEAD não for aceito), com timeout configurável e concorrência limitada; sem a flag, MUST ser ignoradas.
- **FR-2** Respostas 2xx e 3xx MUST ser aceitas; 4xx, 5xx e falhas de rede MUST ser reportadas com o status.
- **FR-3 (D3)** `--format json` MUST imprimir uma lista de objetos com `file`, `line`, `target` e `reason`.
- **FR-4** Um arquivo `.linkcheck.yml` MAY listar padrões de destino a ignorar.

## Critérios de aceite

- **AC-1** Dado um link para um servidor de teste que responde 404, quando `linkcheck --external` roda, então aponta o 404 e sai com 1.
- **AC-2** Dado o mesmo link, quando `linkcheck` roda sem `--external`, então sai com 0.
- **AC-3** Dado um link quebrado, quando `linkcheck --format json` roda, então stdout é um JSON válido com o problema.
- **AC-4** Dado `ignore: ["https://exemplo.invalid/*"]` no `.linkcheck.yml`, quando `linkcheck --external` roda, então o destino é ignorado.

## Decisões

- D1 e D3 respondidas pelo dono em 2026-10-02 conforme as recomendações de [ANALYSIS.md §7](../ANALYSIS.md#7-decisões-respondidas-pelo-dono-em-2026-10-02).
