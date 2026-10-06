# linkcheck

[![CI](https://github.com/fabiodrneles/sdd-kit-demo/actions/workflows/ci.yml/badge.svg)](https://github.com/fabiodrneles/sdd-kit-demo/actions/workflows/ci.yml)

Encontra **links quebrados em arquivos Markdown**: links relativos para arquivos que não existem e âncoras (`#secao`) que não existem no arquivo de destino. Feito para rodar no CI de qualquer repositório com documentação.

Este é também o repositório de demonstração do [sdd-kit](https://github.com/fabiodrneles/sdd-kit): o projeto foi construído do zero com Spec Driven Development, e cada passo está registrado (veja [Como este projeto foi feito](#como-este-projeto-foi-feito)).

## Instalação

```text
go install github.com/fabiodrneles/sdd-kit-demo@latest
```

O binário instalado se chama `sdd-kit-demo`; os exemplos abaixo usam o nome `linkcheck`, como no `go build -o linkcheck`.

## Uso

```bash
# Verifica todos os .md a partir do diretório atual
linkcheck

# Só alguns caminhos
linkcheck docs README.md

# Também as URLs http(s), com até 5 s por requisição
linkcheck --external --timeout 5s

# Saída em JSON (lista de objetos com file, line, target e reason)
linkcheck --format json

# Versão
linkcheck --version
```

Cada problema sai numa linha no formato `arquivo:linha: motivo: destino`, que os editores reconhecem como link:

```text
docs/guia.md:12: arquivo não encontrado: ../imagens/fluxo.png
README.md:40: âncora não encontrada: #instalacao
```

| Código de saída | Significado |
|---|---|
| `0` | nenhum link quebrado |
| `1` | há links quebrados |
| `2` | erro de uso (flag desconhecida, caminho inexistente) |

### O que é verificado

- Links inline, imagens e definições de referência, fora de blocos de código.
- Destinos relativos ao arquivo, e `/caminho` a partir da raiz.
- Âncoras com os slugs do GitHub: `## Instalação` vira `#instalação`, e títulos repetidos ganham `-1`, `-2`.
- URLs externas (`https://...`) ficam de fora por padrão. Com `--external`, cada URL é pedida uma vez (HEAD, ou GET se o servidor recusar HEAD), com até 8 requisições ao mesmo tempo: 2xx e 3xx passam, 4xx e 5xx saem como `HTTP 404`, e falhas como `erro de rede`.

## Como este projeto foi feito

1. [`specs/ANALYSIS.md`](specs/ANALYSIS.md): decisões iniciais do dono (D1–D4).
2. [`specs/`](specs/README.md): constituição, specs com critérios de aceite e o [ROADMAP](specs/ROADMAP.md).
3. [Issues](https://github.com/fabiodrneles/sdd-kit-demo/issues?q=is%3Aissue): um épico por fase, um ticket por tarefa.
4. [Pull requests](https://github.com/fabiodrneles/sdd-kit-demo/pulls?q=is%3Apr): um por ticket, com CI verde e cada critério de aceite citado por um teste.

## Licença

[MIT](LICENSE)
