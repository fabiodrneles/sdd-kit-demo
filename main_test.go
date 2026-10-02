package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// tree writes files (path → content) in a temp dir and chdirs into it.
func tree(t *testing.T, files map[string]string) {
	t.Helper()
	dir := t.TempDir()
	for name, content := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	t.Chdir(dir)
}

func runCLI(args ...string) (string, string, int) {
	var out, errOut strings.Builder
	code := run(args, &out, &errOut)
	return out.String(), errOut.String(), code
}

// 001 AC-1
func TestMissingFile(t *testing.T) {
	tree(t, map[string]string{"a.md": "[x](b.md)\n"})
	out, _, code := runCLI()
	if out != "a.md:1: arquivo não encontrado: b.md\n" || code != 1 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 001 AC-2
func TestMissingAnchor(t *testing.T) {
	tree(t, map[string]string{"a.md": "# A\n\n[x](#nao-existe)\n"})
	out, _, code := runCLI()
	if out != "a.md:3: âncora não encontrada: #nao-existe\n" || code != 1 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 001 AC-3
func TestAnchorInOtherFile(t *testing.T) {
	tree(t, map[string]string{
		"a.md":                "[x](docs/b.md#instalação) [y](docs/b.md#instala%C3%A7%C3%A3o)\n",
		"docs/b.md":           "# B\n\n## Instalação\n",
		"docs/img.png":        "png",
		"docs/guia.md":        "![i](img.png) [a](../a.md) [ext](https://exemplo.invalid/x) [m](mailto:a@b)\n",
		".git/x.md":           "[quebrado](nada.md)\n",
		"node_modules/y/z.md": "[quebrado](nada.md)\n",
	})
	out, _, code := runCLI()
	if out != "" || code != 0 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 001 AC-5
func TestAllValid(t *testing.T) {
	tree(t, map[string]string{"a.md": "# A\n\n[b](b.md) [topo](#a)\n", "b.md": "# B\n"})
	out, errOut, code := runCLI(".")
	if out != "" || errOut != "" || code != 0 {
		t.Errorf("saída %q, erro %q, código %d", out, errOut, code)
	}
}

// 001 AC-6
func TestUnknownFlag(t *testing.T) {
	tree(t, map[string]string{"a.md": "ok\n"})
	if _, _, code := runCLI("--nao-existe"); code != 2 {
		t.Errorf("código %d, quero 2", code)
	}
	if _, _, code := runCLI("nao/existe"); code != 2 {
		t.Errorf("caminho inexistente: código %d, quero 2", code)
	}
}

// 001 FR-5: problems sorted by file and line.
func TestOrder(t *testing.T) {
	tree(t, map[string]string{"b.md": "[x](1.md)\n", "a.md": "\n[x](2.md)\n[x](3.md)\n"})
	out, _, _ := runCLI()
	want := "a.md:2: arquivo não encontrado: 2.md\na.md:3: arquivo não encontrado: 3.md\nb.md:1: arquivo não encontrado: 1.md\n"
	if out != want {
		t.Errorf("saída\n%s\nquero\n%s", out, want)
	}
}
