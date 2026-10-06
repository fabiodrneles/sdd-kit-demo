// Package check finds broken local links and anchors in Markdown files.
package check

import (
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/fabiodrneles/sdd-kit-demo/internal/markdown"
)

// Problem is a broken link.
type Problem struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Reason string `json:"reason"`
	Target string `json:"target"`
}

func (p Problem) String() string {
	return fmt.Sprintf("%s:%d: %s: %s", p.File, p.Line, p.Reason, p.Target)
}

const (
	ReasonMissingFile   = "arquivo não encontrado"
	ReasonMissingAnchor = "âncora não encontrada"
)

var scheme = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9+.-]*:`)

var skipDirs = map[string]bool{".git": true, "node_modules": true}

// Files returns the Markdown files under paths, sorted (spec 001 FR-1).
func Files(paths []string) ([]string, error) {
	var files []string
	for _, root := range paths {
		info, err := os.Stat(root)
		if err != nil {
			return nil, err
		}
		if !info.IsDir() {
			files = append(files, root)
			continue
		}
		err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() && path != root && skipDirs[d.Name()] {
				return filepath.SkipDir
			}
			if !d.IsDir() && strings.EqualFold(filepath.Ext(path), ".md") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Strings(files)
	return files, nil
}

// Local checks the relative links and anchors of each file (spec 001 FR-3,
// FR-4) and returns the problems sorted by file and line (FR-5).
func Local(files []string) ([]Problem, error) {
	anchors := map[string]map[string]bool{}
	anchorsOf := func(path string) (map[string]bool, error) {
		if a, ok := anchors[path]; ok {
			return a, nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		a := markdown.Anchors(string(data))
		anchors[path] = a
		return a, nil
	}

	var problems []Problem
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		for _, link := range markdown.Links(string(data)) {
			if link.Target == "" || scheme.MatchString(link.Target) {
				continue // external or other schemes (spec 002)
			}
			rawPath, fragment, _ := strings.Cut(link.Target, "#")
			p, err := url.PathUnescape(rawPath)
			if err != nil {
				p = rawPath
			}
			dest := file
			if p != "" {
				if strings.HasPrefix(p, "/") {
					dest = filepath.Clean(strings.TrimPrefix(p, "/"))
				} else {
					dest = filepath.Join(filepath.Dir(file), p)
				}
				if _, err := os.Stat(dest); err != nil {
					problems = append(problems, Problem{file, link.Line, ReasonMissingFile, link.Target})
					continue
				}
			}
			if fragment == "" || !strings.EqualFold(filepath.Ext(dest), ".md") {
				continue
			}
			a, err := anchorsOf(dest)
			if err != nil {
				return nil, err
			}
			frag, err := url.PathUnescape(fragment)
			if err != nil {
				frag = fragment
			}
			if !a[strings.ToLower(frag)] {
				problems = append(problems, Problem{file, link.Line, ReasonMissingAnchor, link.Target})
			}
		}
	}
	Sort(problems)
	return problems, nil
}

// Sort orders problems by file and line (spec 001 FR-5).
func Sort(problems []Problem) {
	sort.SliceStable(problems, func(i, j int) bool {
		if problems[i].File != problems[j].File {
			return problems[i].File < problems[j].File
		}
		return problems[i].Line < problems[j].Line
	})
}
