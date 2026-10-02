// Command linkcheck finds broken links in Markdown files (specs/001).
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "linkcheck: em construção; veja specs/ROADMAP.md")
	os.Exit(2)
}
