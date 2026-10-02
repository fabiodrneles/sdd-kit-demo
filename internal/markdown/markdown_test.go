package markdown

import (
	"reflect"
	"testing"
)

func TestLinks(t *testing.T) {
	doc := "# Título\n" +
		"Veja [a](a.md) e ![img](img/x.png \"t\").\n" +
		"[ref]: <docs/guia.md>\n" +
		"Texto `[não](codigo.md)` e [b](<b c.md>) [c](c.md#sec)\n"
	want := []Link{{2, "a.md"}, {2, "img/x.png"}, {3, "docs/guia.md"}, {4, "b c.md"}, {4, "c.md#sec"}}
	if got := Links(doc); !reflect.DeepEqual(got, want) {
		t.Errorf("Links = %v\nquero %v", got, want)
	}
}

// 001 AC-4: links inside fenced code blocks are ignored.
func TestLinksIgnoreFencedCode(t *testing.T) {
	doc := "antes [a](a.md)\n```md\n[quebrado](nao.md)\n~~~\n```\n~~~\n[x](y.md)\n~~~\ndepois [b](b.md)\n"
	want := []Link{{1, "a.md"}, {9, "b.md"}}
	if got := Links(doc); !reflect.DeepEqual(got, want) {
		t.Errorf("Links = %v\nquero %v", got, want)
	}
}

func TestAnchors(t *testing.T) {
	doc := "# Instalação\n## Uso: `linkcheck`!\n## Uso: `linkcheck`!\n```\n# não é título\n```\n### C# e .NET\n"
	got := Anchors(doc)
	for _, a := range []string{"instalação", "uso-linkcheck", "uso-linkcheck-1", "c-e-net"} {
		if !got[a] {
			t.Errorf("âncora %q ausente em %v", a, got)
		}
	}
	if got["não-é-título"] {
		t.Error("título dentro de bloco de código virou âncora")
	}
}
