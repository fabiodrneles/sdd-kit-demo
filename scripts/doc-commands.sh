#!/bin/sh
# Runs every ```bash block of README.md in a sample docs tree with the freshly
# built linkcheck on PATH. Commands that cannot run here belong in ```text.
set -eu
root=$(pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
go build -o "$tmp/bin/linkcheck" .
mkdir -p "$tmp/sample/docs"
printf '# Exemplo\n\nVeja o [guia](docs/guia.md#uso).\n' > "$tmp/sample/README.md"
printf '# Guia\n\n## Uso\n\nVolte ao [início](../README.md).\n' > "$tmp/sample/docs/guia.md"
blocks=$(awk '/^```bash$/{b=1; n++; next} /^```$/{b=0} b{print > ("'"$tmp"'/block" n ".sh")} END{print n+0}' "$root/README.md")
i=1
while [ "$i" -le "$blocks" ]; do
	echo "== README bloco bash $i"
	(cd "$tmp/sample" && PATH="$tmp/bin:$PATH" sh -eu "$tmp/block$i.sh")
	i=$((i + 1))
done
echo "doc-commands: $blocks bloco(s) ok"
