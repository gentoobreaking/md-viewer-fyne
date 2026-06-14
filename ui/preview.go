package ui

import (
	"md-viewer-app/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

// previewCtrl manages the markdown preview area.
// It renders markdown using goldmark (full GFM) and displays via embedded HTTP server + browser.
type previewCtrl struct {
	server     *core.Server
	status     *widget.Label
	scroll     *container.Scroll
	currentURL string
}

func newPreview() *previewCtrl {
	p := &previewCtrl{}

	// Start embedded HTTP server
	srv, err := core.NewServer()
	if err != nil {
		p.status = widget.NewLabel("⚠️ 預覽伺服器啟動失敗")
		p.scroll = container.NewScroll(p.status)
		return p
	}
	p.server = srv

	// In-app status display
	p.status = widget.NewLabel("在瀏覽器中預覽 Markdown\n\n支援 GFM 語法：表格、任務清單、刪除線\n\n使用 ⌘O 開啟檔案")
	p.status.Wrapping = fyne.TextWrapWord
	p.status.Alignment = fyne.TextAlignCenter

	p.scroll = container.NewVScroll(p.status)

	return p
}

// widget returns the preview area's fyne CanvasObject.
func (p *previewCtrl) widget() fyne.CanvasObject {
	return p.scroll
}

// show renders markdown content and opens it in the browser.
func (p *previewCtrl) show(content string, filePath string) {
	// Parse markdown to HTML
	htmlContent, err := core.ParseMarkdown([]byte(content))
	if err != nil {
		p.status.SetText("⚠️ Markdown 解析失敗\n\n" + err.Error())
		return
	}

	// Extract title and render full HTML page
	title := core.ExtractTitle([]byte(content), core.BaseName(filePath))
	fullHTML := core.RenderHTML(title, htmlContent, filePath)

	// Update the embedded server
	if p.server != nil {
		p.server.SetHTML(fullHTML)
		p.currentURL = p.server.PreviewURL()
	}

	// Update in-app status
	p.status.SetText("📄 " + core.BaseName(filePath) + "\n\n預覽已更新")

	// Open in browser
	p.openInBrowser()
}

// openInBrowser opens the preview URL in the default browser.
func (p *previewCtrl) openInBrowser() {
	if p.currentURL == "" {
		return
	}
	u, err := url.Parse(p.currentURL)
	if err != nil {
		return
	}
	fyne.CurrentApp().OpenURL(u)
}

// close shuts down the preview server.
func (p *previewCtrl) close() {
	if p.server != nil {
		p.server.Close()
	}
}
