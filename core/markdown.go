package core

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// ParseMarkdown converts markdown content to HTML using goldmark with full GFM support.
func ParseMarkdown(content []byte) (string, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
			extension.TaskList,
			extension.DefinitionList,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ExtractTitle extracts the first heading from markdown content.
// Falls back to the file name if no heading is found.
func ExtractTitle(content []byte, filename string) string {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			title := strings.TrimSpace(strings.TrimPrefix(trimmed, "#"))
			if title != "" {
				return title
			}
		}
	}
	return filename
}
