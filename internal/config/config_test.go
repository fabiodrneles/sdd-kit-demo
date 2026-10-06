package config

import (
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tc := range []struct {
		name, data string
		want       []string
	}{
		{"inline", `ignore: ["https://a/*", 'b.md']`, []string{"https://a/*", "b.md"}},
		{"bloco", "ignore:\n  - https://a/*\n  - \"c # d\"\n", []string{"https://a/*", "c # d"}},
		{"comentário e vazio", "# nada\n\nignore: [] # vazio\n", nil},
	} {
		cfg, err := Parse(tc.data, "x")
		if err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
		if len(cfg.Ignore) != len(tc.want) {
			t.Fatalf("%s: %q, esperava %q", tc.name, cfg.Ignore, tc.want)
		}
		for i := range tc.want {
			if cfg.Ignore[i] != tc.want[i] {
				t.Errorf("%s: %q, esperava %q", tc.name, cfg.Ignore, tc.want)
			}
		}
	}
	for _, bad := range []string{"outra: [a]", "ignore: a", "  ignore: [a]", "sem dois pontos"} {
		if _, err := Parse(bad, "x"); err == nil {
			t.Errorf("%q devia dar erro", bad)
		}
	}
}

func TestLoadMissing(t *testing.T) {
	cfg, err := Load(filepath.Join(t.TempDir(), File))
	if err != nil || len(cfg.Ignore) != 0 {
		t.Errorf("arquivo ausente: %+v, %v", cfg, err)
	}
}

func TestIgnored(t *testing.T) {
	cfg := Config{Ignore: []string{"https://exemplo.invalid/*", "docs/*.md", "exato.md", "*x*y"}}
	for target, want := range map[string]bool{
		"https://exemplo.invalid/a/b": true,
		"https://exemplo.invalid":     false,
		"docs/a/b.md":                 true,
		"docs/a.txt":                  false,
		"exato.md":                    true,
		"exato.md#x":                  false,
		"axbxy":                       true,
		"ay":                          false,
	} {
		if got := cfg.Ignored(target); got != want {
			t.Errorf("Ignored(%q) = %v, esperava %v", target, got, want)
		}
	}
}
