# Changelog

Todas as mudanças relevantes deste projeto. Formato [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/); versões [SemVer](https://semver.org/lang/pt-BR/).

## [Unreleased]

## [0.1.0] - 2026-10-02

### Adicionado

- `linkcheck [CAMINHO...]`: encontra links relativos para arquivos que não existem e âncoras que não existem, em links inline, imagens e definições de referência, fora de blocos de código.
- Âncoras com os slugs do GitHub, inclusive acentos e títulos repetidos.
- Saída `arquivo:linha: motivo: destino` e códigos de saída 0, 1 e 2; `--version`.
- README com os comandos verificados no CI; o próprio `linkcheck` verifica a documentação do repositório.
