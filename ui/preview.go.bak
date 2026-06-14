package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// previewCtrl manages the markdown preview area
type previewCtrl struct {
	richText *widget.RichText
	scroll   *container.Scroll
}

func newPreview() *previewCtrl {
	p := &previewCtrl{}

	p.richText = widget.NewRichTextFromMarkdown(
		"# md-viewer\n\n歡迎使用 md-viewer！\n\n請使用 **⌘O** 開啟檔案，或 **⌘⇧O** 開啟資料夾。")

	p.scroll = container.NewScroll(p.richText)

	return p
}

// widget returns the preview area's fyne CanvasObject
func (p *previewCtrl) widget() fyne.CanvasObject {
	return p.scroll
}

// show renders markdown content in the preview area
func (p *previewCtrl) show(content string) {
	p.richText.ParseMarkdown(content)
	p.richText.Refresh()
	p.scroll.ScrollToTop()
}
