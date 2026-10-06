// Package config reads .linkcheck.yml (spec 002 FR-4). Only the standard
// library is used, so it understands the small YAML subset the file needs:
//
//	ignore: ["https://exemplo.invalid/*", "docs/old.md"]
//
// or the same list in block form ("ignore:" followed by "- pattern" lines).
// Comments (#) and blank lines are allowed; any other key is an error, so a
// typo never silently disables the ignore list.
package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// File is the default configuration file name, read from the current directory.
const File = ".linkcheck.yml"

// Config is the parsed configuration.
type Config struct {
	Ignore []string // target patterns; "*" matches any run of characters
}

// Load reads path. A missing file is not an error: it yields an empty Config.
func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return Config{}, nil
	}
	if err != nil {
		return Config{}, err
	}
	return Parse(string(data), path)
}

// Parse parses the contents of a configuration file named name (used in errors).
func Parse(data, name string) (Config, error) {
	var cfg Config
	inList := false
	sc := bufio.NewScanner(strings.NewReader(data))
	for n := 1; sc.Scan(); n++ {
		line := stripComment(sc.Text())
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if inList && strings.HasPrefix(trimmed, "- ") {
			cfg.Ignore = append(cfg.Ignore, unquote(strings.TrimSpace(trimmed[2:])))
			continue
		}
		inList = false
		key, value, ok := strings.Cut(trimmed, ":")
		if !ok || line != strings.TrimLeft(line, " \t") {
			return Config{}, fmt.Errorf("%s:%d: linha inválida: %q", name, n, trimmed)
		}
		if strings.TrimSpace(key) != "ignore" {
			return Config{}, fmt.Errorf("%s:%d: chave desconhecida %q (só existe ignore)", name, n, strings.TrimSpace(key))
		}
		value = strings.TrimSpace(value)
		switch {
		case value == "":
			inList = true
		case strings.HasPrefix(value, "[") && strings.HasSuffix(value, "]"):
			for _, item := range strings.Split(value[1:len(value)-1], ",") {
				if item = strings.TrimSpace(item); item != "" {
					cfg.Ignore = append(cfg.Ignore, unquote(item))
				}
			}
		default:
			return Config{}, fmt.Errorf("%s:%d: ignore deve ser uma lista", name, n)
		}
	}
	return cfg, sc.Err()
}

// Ignored reports whether target matches one of the ignore patterns.
func (c Config) Ignored(target string) bool {
	for _, p := range c.Ignore {
		if match(p, target) {
			return true
		}
	}
	return false
}

// match is a glob where "*" matches any run of characters, "/" included.
func match(pattern, s string) bool {
	parts := strings.Split(pattern, "*")
	if len(parts) == 1 {
		return pattern == s
	}
	if !strings.HasPrefix(s, parts[0]) {
		return false
	}
	s = s[len(parts[0]):]
	for _, mid := range parts[1 : len(parts)-1] {
		i := strings.Index(s, mid)
		if i < 0 {
			return false
		}
		s = s[i+len(mid):]
	}
	return strings.HasSuffix(s, parts[len(parts)-1])
}

func stripComment(line string) string {
	inQuote := byte(0)
	for i := 0; i < len(line); i++ {
		switch c := line[i]; {
		case inQuote != 0 && c == inQuote:
			inQuote = 0
		case inQuote == 0 && (c == '"' || c == '\''):
			inQuote = c
		case inQuote == 0 && c == '#':
			return line[:i]
		}
	}
	return line
}

func unquote(s string) string {
	if len(s) >= 2 && (s[0] == '"' || s[0] == '\'') && s[len(s)-1] == s[0] {
		return s[1 : len(s)-1]
	}
	return s
}
