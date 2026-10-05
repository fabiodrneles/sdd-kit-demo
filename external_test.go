package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// server answers /ok 200, /missing 404, /moved 301 and /nohead 405 to HEAD
// (200 to GET).
func server(t *testing.T) string {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/ok":
		case "/moved":
			http.Redirect(w, r, "/ok", http.StatusMovedPermanently)
		case "/nohead":
			if r.Method == http.MethodHead {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(srv.Close)
	return srv.URL
}

// 002 AC-1
func TestExternalNotFound(t *testing.T) {
	url := server(t)
	tree(t, map[string]string{"a.md": "[x](" + url + "/missing)\n"})
	out, _, code := runCLI("--external")
	if out != "a.md:1: HTTP 404: "+url+"/missing\n" || code != 1 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 002 AC-2
func TestExternalIgnoredWithoutFlag(t *testing.T) {
	url := server(t)
	tree(t, map[string]string{"a.md": "[x](" + url + "/missing)\n"})
	if out, _, code := runCLI(); out != "" || code != 0 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 002 FR-2: 2xx e 3xx passam, e um servidor que recusa HEAD é verificado com GET.
func TestExternalAccepted(t *testing.T) {
	url := server(t)
	tree(t, map[string]string{"a.md": "[a](" + url + "/ok) [b](" + url + "/moved)\n\n[c](" + url + "/nohead)\n"})
	if out, _, code := runCLI("--external"); out != "" || code != 0 {
		t.Errorf("saída %q, código %d", out, code)
	}
}

// 002 FR-2: falha de rede é reportada.
func TestExternalNetworkError(t *testing.T) {
	url := server(t)
	tree(t, map[string]string{"a.md": "[x](http://127.0.0.1:1/)\n[y](" + url + "/ok)\n"})
	out, _, code := runCLI("--external", "--timeout", "2s")
	if !strings.HasPrefix(out, "a.md:1: erro de rede: ") || strings.Count(out, "\n") != 1 || code != 1 {
		t.Errorf("saída %q, código %d", out, code)
	}
}
