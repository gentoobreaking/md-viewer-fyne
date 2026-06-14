package core

import (
	"fmt"
	"strings"
)

// RenderHTML wraps rendered markdown HTML in a full page with GitHub-like CSS styling.
func RenderHTML(title, htmlContent, filePath string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-TW">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>%s</title>
  <style>%s</style>
</head>
<body>
  <div class="meta">&#128196; %s</div>
  <article class="markdown-body">
%s
  </article>
  <script>%s</script>
</body>
</html>`, escapeAttr(title), cssStyles, escapeHTML(filePath), htmlContent, jsCode)
}

func escapeAttr(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, `"`, "&quot;")
	return s
}

func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

const cssStyles = `
/* Reset & Base */
*, *::before, *::after { box-sizing: border-box; }
body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji";
  font-size: 16px;
  line-height: 1.7;
  max-width: 860px;
  margin: 0 auto;
  padding: 40px 32px;
  color: #1f2328;
  background: #ffffff;
}

/* Meta bar */
.meta {
  font-size: 13px;
  color: #656d76;
  margin-bottom: 24px;
  padding: 8px 12px;
  background: #f6f8fa;
  border-radius: 6px;
  border-left: 4px solid #0969da;
}

/* Markdown body */
.markdown-body { }
.markdown-body > *:first-child { margin-top: 0 !important; }

/* Headings */
h1, h2, h3, h4, h5, h6 {
  margin-top: 24px;
  margin-bottom: 16px;
  font-weight: 600;
  line-height: 1.25;
}
h1 { font-size: 2em; border-bottom: 1px solid #d1d9e0; padding-bottom: 0.3em; }
h2 { font-size: 1.5em; border-bottom: 1px solid #d1d9e0; padding-bottom: 0.3em; }
h3 { font-size: 1.25em; }
h4 { font-size: 1em; }

/* Paragraphs */
p { margin: 0 0 16px; }

/* Links */
a { color: #0969da; text-decoration: none; }
a:hover { text-decoration: underline; }

/* Inline code */
code {
  font-family: "SFMono-Regular", Consolas, "Liberation Mono", Menlo, monospace;
  font-size: 85%;
  background: #eff1f3;
  padding: 0.2em 0.4em;
  border-radius: 6px;
}

/* Code blocks */
pre {
  background: #f6f8fa;
  padding: 16px;
  border-radius: 6px;
  overflow-x: auto;
  line-height: 1.45;
  margin: 0 0 16px;
  border: 1px solid #d1d9e0;
}
pre code {
  background: transparent;
  padding: 0;
  font-size: 100%;
  border-radius: 0;
}

/* Blockquotes */
blockquote {
  margin: 0 0 16px;
  padding: 0 1em;
  color: #656d76;
  border-left: 0.25em solid #d1d9e0;
}
blockquote p:last-child { margin-bottom: 0; }

/* Lists */
ul, ol { margin: 0 0 16px; padding-left: 2em; }
li { margin-bottom: 4px; }
li > p { margin-bottom: 8px; }

/* Task lists */
.task-list-item { list-style: none; margin-left: -1.5em; }
.task-list-item input[type="checkbox"] {
  margin-right: 0.5em;
  accent-color: #0969da;
}

/* Tables */
table {
  border-collapse: collapse;
  width: 100%;
  margin-bottom: 16px;
  overflow: auto;
  display: block;
}
th, td {
  border: 1px solid #d1d9e0;
  padding: 6px 13px;
}
th {
  background: #f6f8fa;
  font-weight: 600;
}
tr:nth-child(even) {
  background: #f6f8fa;
}

/* Horizontal rule */
hr {
  border: none;
  border-top: 1px solid #d1d9e0;
  margin: 24px 0;
}

/* Images */
img {
  max-width: 100%;
  height: auto;
  border-radius: 6px;
}

/* Strikethrough */
del {
  color: #656d76;
}

/* Definition lists */
dl { margin: 0 0 16px; }
dt { font-weight: 600; margin-top: 8px; }
dd { margin-left: 2em; }

/* Scrollbar */
::-webkit-scrollbar { width: 8px; height: 8px; }
::-webkit-scrollbar-track { background: transparent; }
::-webkit-scrollbar-thumb { background: #d1d9e0; border-radius: 4px; }
::-webkit-scrollbar-thumb:hover { background: #a1a9b1; }

/* ===== Dark Mode ===== */
@media (prefers-color-scheme: dark) {
  body {
    background: #0d1117;
    color: #e6edf3;
  }
  .meta {
    background: #161b22;
    color: #8b949e;
    border-left-color: #58a6ff;
  }
  h1, h2 { border-color: #30363d; }
  code {
    background: #30363d;
    color: #e6edf3;
  }
  pre {
    background: #161b22;
    border-color: #30363d;
    color: #e6edf3;
  }
  blockquote {
    color: #8b949e;
    border-color: #30363d;
  }
  a { color: #58a6ff; }
  table th {
    background: #161b22;
  }
  th, td {
    border-color: #30363d;
  }
  tr:nth-child(even) {
    background: #161b22;
  }
  hr { border-color: #30363d; }
  del { color: #8b949e; }
  ::-webkit-scrollbar-thumb { background: #30363d; }
  ::-webkit-scrollbar-thumb:hover { background: #484f58; }
}
`

const jsCode = `
// Auto-reload support via EventSource
(function() {
  if (window.__mdViewerReload) return;
  window.__mdViewerReload = true;

  var source = new EventSource('/events');
  source.onmessage = function(e) {
    if (e.data === 'reload') {
      location.reload();
    }
  };
  source.onerror = function() {
    // Server disconnected, stop reconnecting after 3 attempts
    if (source.readyState === EventSource.CLOSED) return;
  };
})();
`
