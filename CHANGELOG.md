# Changelog

Todas as mudanças relevantes deste projeto. Formato [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/); versões [SemVer](https://semver.org/lang/pt-BR/).

## [Unreleased]

## [1.0.0] - 2026-10-06

Fase 3 do [ROADMAP](specs/ROADMAP.md): distribuição (spec 003).

### Adicionado

- **Binários prontos:** cada tag publica o `linkcheck` para Linux, macOS e Windows (amd64 e arm64) com `checksums.txt`, pelo GoReleaser, depois de `make ci` verde. A versão é calculada pelo go-release-manager e o *Release check* ensaia a release em todo PR (003 FR-1, FR-2, #28).
- **GitHub Action:** `uses: fabiodrneles/sdd-kit-demo@v1.0.0` roda o `linkcheck` no repositório, com as entradas `paths`, `external`, `timeout` e `format`; o job falha se houver link quebrado (003 FR-4, #29).

## [0.2.0] - 2026-10-06

Fase 2 do [ROADMAP](specs/ROADMAP.md): links externos e integração (spec 002).

### Adicionado

- **`--external`:** verifica também as URLs http(s), com HEAD (ou GET, se o servidor recusar HEAD), uma vez por URL, até 8 em paralelo e com `--timeout` configurável. 2xx e 3xx passam; 4xx/5xx saem como `HTTP <status>` e falhas como `erro de rede` (002 FR-1, FR-2, #18).
- **`--format json`:** a lista de problemas em JSON, com `file`, `line`, `target` e `reason` (`[]` sem problemas), para outras ferramentas consumirem (002 FR-3, #21).
- **`.linkcheck.yml`:** `ignore` lista padrões de destino a ignorar, locais ou externos, e com `--external` uma URL ignorada nem é pedida. Uma chave desconhecida sai com 2, para um erro de digitação não desligar a lista (002 FR-4, #23).

## [0.1.0] - 2026-10-02

### Adicionado

- `linkcheck [CAMINHO...]`: encontra links relativos para arquivos que não existem e âncoras que não existem, em links inline, imagens e definições de referência, fora de blocos de código.
- Âncoras com os slugs do GitHub, inclusive acentos e títulos repetidos.
- Saída `arquivo:linha: motivo: destino` e códigos de saída 0, 1 e 2; `--version`.
- README com os comandos verificados no CI; o próprio `linkcheck` verifica a documentação do repositório.
