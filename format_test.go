package main

import (
	"encoding/json"
	"testing"
)

// 002 AC-3
func TestFormatJSON(t *testing.T) {
	tree(t, map[string]string{"a.md": "[x](missing.md)\n"})
	out, _, code := runCLI("--format", "json")
	var got []struct {
		File   string `json:"file"`
		Line   int    `json:"line"`
		Target string `json:"target"`
		Reason string `json:"reason"`
	}
	if err := json.Unmarshal([]byte(out), &got); err != nil {
		t.Fatalf("stdout não é JSON válido: %v\n%s", err, out)
	}
	if code != 1 || len(got) != 1 || got[0].File != "a.md" || got[0].Line != 1 || got[0].Target != "missing.md" || got[0].Reason == "" {
		t.Errorf("problemas %+v, código %d", got, code)
	}
}

func TestFormatJSONEmpty(t *testing.T) {
	tree(t, map[string]string{"a.md": "sem links\n"})
	out, _, code := runCLI("--format", "json")
	if out != "[]\n" || code != 0 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

func TestFormatInvalid(t *testing.T) {
	tree(t, map[string]string{"a.md": "sem links\n"})
	_, errOut, code := runCLI("--format", "xml")
	if code != 2 || errOut == "" {
		t.Errorf("stderr %q, código %d", errOut, code)
	}
}
