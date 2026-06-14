package ui

import (
	"fmt"
	"math"
	"time"

	"md-viewer-app/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Run starts the application
func Run() {
	a := app.NewWithID("com.mdviewer.app")
	a.Settings().SetTheme(&mdViewerTheme{})

	w := a.NewWindow("md-viewer")
	w.Resize(fyne.NewSize(900, 600))

	// Create components
	sidebar := newSidebar()
	preview := newPreview()
	defer preview.close()

	// Wire sidebar selection to preview
	sidebar.onSelect = func(path string) {
		content, err := core.ReadFile(path)
		if err != nil {
			preview.show("⚠️ 讀取失敗: "+err.Error(), path)
			return
		}
		preview.show(content, path)
		w.SetTitle("md-viewer — " + core.BaseName(path))
	}

	// Layout: sidebar + preview
	split := container.NewHSplit(sidebar.widget(), preview.widget())
	split.SetOffset(0.25)

	w.SetContent(split)

	// Menu bar + keyboard shortcuts
	setupMenu(a, w, sidebar, preview)

	w.SetCloseIntercept(func() {
		preview.close()
		w.Close()
	})

	w.ShowAndRun()
}

// showZoomIndicator shows a temporary zoom percentage overlay
func showZoomIndicator(w fyne.Window, width, height float32) {
	baseW, baseH := float32(900), float32(600)
	zoom := (width/baseW + height/baseH) / 2
	percent := int(zoom * 100)

	label := widget.NewLabel(fmt.Sprintf("%d%%", percent))
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.Alignment = fyne.TextAlignCenter

	cont := fyne.NewContainer(label)
	cont.Resize(fyne.NewSize(60, 30))

	overlay := fyne.NewContainerWithoutLayout(cont)
	overlay.Move(fyne.NewPos(w.Canvas().Size().Width-80, 10))

	w.Canvas().Overlays().Add(overlay)

	go func() {
		time.Sleep(1500 * time.Millisecond)
		w.Canvas().Overlays().Remove(overlay)
	}()
}

// setupMenu configures the application menu bar and keyboard shortcuts
func setupMenu(a fyne.App, w fyne.Window, sidebar *sidebarCtrl, preview *previewCtrl) {
	openFileItem := fyne.NewMenuItem("開啟檔案...", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			path := reader.URI().Path()
			reader.Close()

			content, err := core.ReadFile(path)
			if err != nil {
				preview.show("⚠️ 讀取失敗: "+err.Error(), path)
				return
			}
			preview.show(content, path)
			w.SetTitle("md-viewer — " + core.BaseName(path))
		}, w)
	})

	openFolderItem := fyne.NewMenuItem("開啟資料夾...", func() {
		dialog.ShowFolderOpen(func(lister fyne.ListableURI, err error) {
			if err != nil || lister == nil {
				return
			}
			sidebar.loadDirectory(lister.String())
		}, w)
	})

	toggleSidebarItem := fyne.NewMenuItem("切換側邊欄", func() {
		sidebar.toggle()
	})

	refreshItem := fyne.NewMenuItem("重新整理預覽", func() {
		path := sidebar.selectedPath()
		if path != "" && sidebar.onSelect != nil {
			sidebar.onSelect(path)
		}
	})

	fileMenu := fyne.NewMenu("檔案", openFileItem, openFolderItem)
	viewMenu := fyne.NewMenu("顯示", toggleSidebarItem, refreshItem)
	aboutItem := fyne.NewMenuItem("關於 md-viewer", func() {
		dialog.ShowInformation("關於 md-viewer",
			"md-viewer v0.2.0\n\n純閱讀導向的 Markdown 預覽器\n\n✅ GFM 語法支援：表格、任務清單、刪除線\n✅ 程式碼區塊\n✅ 深色模式\n✅ 自動重整\n\nBuilt with Go + goldmark + Fyne",
			w)
	})
	helpMenu := fyne.NewMenu("說明", aboutItem)

	w.SetMainMenu(fyne.NewMainMenu(fileMenu, viewMenu, helpMenu))

	canvas := w.Canvas()

	// --- Method 1: AddShortcut (works on fyne.Canvas directly) ---
	canvas.AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyEqual,
		Modifier: fyne.KeyModifierSuper,
	}, func(_ fyne.Shortcut) {
		size := canvas.Size()
		newW := float32(math.Min(float64(size.Width)*1.1, 2400))
		newH := float32(math.Min(float64(size.Height)*1.1, 1600))
		w.Resize(fyne.NewSize(newW, newH))
		showZoomIndicator(w, newW, newH)
	})
	canvas.AddShortcut(&desktop.CustomShortcut{
		KeyName:  fyne.KeyMinus,
		Modifier: fyne.KeyModifierSuper,
	}, func(_ fyne.Shortcut) {
		size := canvas.Size()
		newW := float32(math.Max(float64(size.Width)*0.9, 600))
		newH := float32(math.Max(float64(size.Height)*0.9, 400))
		w.Resize(fyne.NewSize(newW, newH))
		showZoomIndicator(w, newW, newH)
	})

	// --- Method 2: SetOnTypedKey (fallback, always available on fyne.Canvas) ---
	canvas.SetOnTypedKey(func(key *fyne.KeyEvent) {
		// Cmd+B: toggle sidebar
		if key.Name == fyne.KeyB {
			sidebar.toggle()
			return
		}
		// Cmd+R: refresh
		if key.Name == fyne.KeyR {
			refreshItem.Action()
			return
		}
		// Zoom in: Cmd+= or Cmd++
		if key.Name == fyne.KeyEqual || key.Name == fyne.KeyPlus {
			size := canvas.Size()
			newW := float32(math.Min(float64(size.Width)*1.1, 2400))
			newH := float32(math.Min(float64(size.Height)*1.1, 1600))
			w.Resize(fyne.NewSize(newW, newH))
			showZoomIndicator(w, newW, newH)
			return
		}
		// Zoom out: Cmd+-
		if key.Name == fyne.KeyMinus {
			size := canvas.Size()
			newW := float32(math.Max(float64(size.Width)*0.9, 600))
			newH := float32(math.Max(float64(size.Height)*0.9, 400))
			w.Resize(fyne.NewSize(newW, newH))
			showZoomIndicator(w, newW, newH)
			return
		}
	})
}
