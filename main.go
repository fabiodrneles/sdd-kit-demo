// Command linkcheck finds broken links in Markdown files (specs/001).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fabiodrneles/sdd-kit-demo/internal/check"
)

// Set by GoReleaser through -ldflags "-X main.version=...".
var version = "dev"

// Exit codes (spec 001 FR-6).
const (
	exitOK     = 0
	exitBroken = 1
	exitUsage  = 2
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("linkcheck", flag.ContinueOnError)
	fs.SetOutput(stderr)
	showVersion := fs.Bool("version", false, "imprime a versão e sai")
	external := fs.Bool("external", false, "verifica também as URLs http(s) (spec 002)")
	timeout := fs.Duration("timeout", 10*time.Second, "tempo máximo de cada requisição com --external")
	format := fs.String("format", "text", "formato da saída: text ou json (spec 002)")
	fs.Usage = func() {
		_, _ = fmt.Fprintln(stderr, "uso: linkcheck [flags] [CAMINHO...]")
		_, _ = fmt.Fprintln(stderr, "Verifica links locais e âncoras dos arquivos .md (padrão: diretório atual).")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return exitOK
		}
		return exitUsage
	}
	if *format != "text" && *format != "json" {
		_, _ = fmt.Fprintf(stderr, "linkcheck: --format inválido: %q (use text ou json)\n", *format)
		return exitUsage
	}
	if *showVersion {
		_, _ = fmt.Fprintln(stdout, version)
		return exitOK
	}
	paths := fs.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}
	files, err := check.Files(paths)
	if err != nil {
		_, _ = fmt.Fprintln(stderr, "linkcheck:", err)
		return exitUsage
	}
	problems, err := check.Local(files)
	if err != nil {
		_, _ = fmt.Fprintln(stderr, "linkcheck:", err)
		return exitUsage
	}
	if *external {
		ext, err := check.External(files, check.ExternalOptions{Timeout: *timeout})
		if err != nil {
			_, _ = fmt.Fprintln(stderr, "linkcheck:", err)
			return exitUsage
		}
		problems = append(problems, ext...)
		check.Sort(problems)
	}
	if *format == "json" {
		if problems == nil {
			problems = []check.Problem{}
		}
		enc := json.NewEncoder(stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(problems); err != nil {
			_, _ = fmt.Fprintln(stderr, "linkcheck:", err)
			return exitUsage
		}
	} else {
		for _, p := range problems {
			_, _ = fmt.Fprintln(stdout, p)
		}
	}
	if len(problems) > 0 {
		return exitBroken
	}
	return exitOK
}
