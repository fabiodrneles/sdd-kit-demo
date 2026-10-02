// Package markdown extracts links and heading anchors from Markdown files.
package markdown

import (
	"regexp"
	"strings"
	"unicode"
)

// Link is a link destination and the 1-based line where it appears.
type Link struct {
	Line   int
	Target string
}

var (
	inlineLink = regexp.MustCompile(`!?\[[^\]]*\]\(\s*(<[^>]*>|[^)\s]+)(?:\s+(?:"[^"]*"|'[^']*'))?\s*\)`)
	refDef     = regexp.MustCompile(`^ {0,3}\[[^\]]+\]:\s*(<[^>]*>|\S+)`)
	codeSpan   = regexp.MustCompile("`+[^`]*`+")
	atxHeading = regexp.MustCompile(`^ {0,3}#{1,6}\s+(.*?)\s*#*\s*$`)
	fence      = regexp.MustCompile("^ {0,3}(```+|~~~+)")
)

// lines yields each line outside fenced code blocks with its number.
func lines(content string, fn func(n int, line string)) {
	open := ""
	for i, line := range strings.Split(content, "\n") {
		line = strings.TrimSuffix(line, "\r")
		if m := fence.FindStringSubmatch(line); m != nil {
			switch marker := m[1][:3]; open {
			case "":
				open = marker
				continue
			case marker:
				open = ""
				continue
			}
		}
		if open == "" {
			fn(i+1, line)
		}
	}
}

// Links returns every link destination outside code (spec 001 FR-2).
func Links(content string) []Link {
	var out []Link
	lines(content, func(n int, line string) {
		line = codeSpan.ReplaceAllString(line, "")
		if m := refDef.FindStringSubmatch(line); m != nil {
			out = append(out, Link{Line: n, Target: unwrap(m[1])})
			return
		}
		for _, m := range inlineLink.FindAllStringSubmatch(line, -1) {
			out = append(out, Link{Line: n, Target: unwrap(m[1])})
		}
	})
	return out
}

func unwrap(target string) string {
	return strings.TrimSuffix(strings.TrimPrefix(target, "<"), ">")
}

// Anchors returns the set of heading anchors of a document, using the slugs
// GitHub generates: lowercase, punctuation removed, spaces as hyphens, and a
// numeric suffix for repeated headings (spec 001 FR-4).
func Anchors(content string) map[string]bool {
	anchors := map[string]bool{}
	seen := map[string]int{}
	lines(content, func(_ int, line string) {
		m := atxHeading.FindStringSubmatch(line)
		if m == nil {
			return
		}
		slug := Slug(m[1])
		if n := seen[slug]; n > 0 {
			anchors[slug+"-"+itoa(n)] = true
		} else {
			anchors[slug] = true
		}
		seen[slug]++
	})
	return anchors
}

// Slug converts heading text to its GitHub anchor.
func Slug(text string) string {
	text = codeSpan.ReplaceAllStringFunc(text, func(s string) string { return strings.Trim(s, "`") })
	var b strings.Builder
	for _, r := range strings.ToLower(strings.TrimSpace(text)) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-':
			b.WriteRune(r)
		case r == ' ':
			b.WriteRune('-')
		}
	}
	return b.String()
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var d []byte
	for ; n > 0; n /= 10 {
		d = append([]byte{byte('0' + n%10)}, d...)
	}
	return string(d)
}
