package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// 002 AC-4
func TestConfigIgnoresExternal(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(srv.Close)
	tree(t, map[string]string{
		"a.md":           "[x](" + srv.URL + "/missing)\n",
		".linkcheck.yml": "# links de exemplo\nignore: [\"" + srv.URL + "/*\"]\n",
	})
	out, _, code := runCLI("--external")
	if out != "" || code != 0 || hits.Load() != 0 {
		t.Errorf("saída %q, código %d, requisições %d", out, code, hits.Load())
	}
}

func TestConfigIgnoresLocalBlockList(t *testing.T) {
	tree(t, map[string]string{
		"a.md":           "[x](old.md) [y](gone.md)\n",
		".linkcheck.yml": "ignore:\n  - \"old.md\" # ainda vamos criar\n",
	})
	out, _, code := runCLI()
	if out != "a.md:1: arquivo não encontrado: gone.md\n" || code != 1 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

func TestConfigInvalid(t *testing.T) {
	tree(t, map[string]string{
		"a.md":           "sem links\n",
		".linkcheck.yml": "ignorar: [\"x\"]\n",
	})
	_, errOut, code := runCLI()
	if code != 2 || !strings.Contains(errOut, "chave desconhecida") {
		t.Errorf("stderr %q, código %d", errOut, code)
	}
}
