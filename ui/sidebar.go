package ui

import (
	"md-viewer-app/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// sidebarCtrl manages the sidebar file list
type sidebarCtrl struct {
	list        *widget.List
	files       []core.FileEntry
	hidden      bool
	selectedIdx int
	onSelect    func(path string)
	parent      *fyne.Container
}

// FileEntry holds display info for a sidebar item
type fileEntry = core.FileEntry

func newSidebar() *sidebarCtrl {
	s := &sidebarCtrl{}

	s.list = widget.NewList(
		func() int {
			return len(s.files)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id < len(s.files) {
				label.SetText(s.files[id].Name)
			}
		},
	)

	s.list.OnSelected = func(id widget.ListItemID) {
		s.selectedIdx = int(id)
		if id < len(s.files) && s.onSelect != nil {
			s.onSelect(s.files[id].Path)
		}
	}

	return s
}

// widget returns the sidebar's fyne CanvasObject
func (s *sidebarCtrl) widget() fyne.CanvasObject {
	heading := widget.NewLabel("📂 檔案列表")
	heading.TextStyle = fyne.TextStyle{Bold: true}

	s.parent = container.NewBorder(heading, nil, nil, nil, s.list)
	return s.parent
}

// loadDirectory scans a directory for .md files
func (s *sidebarCtrl) loadDirectory(dirPath string) {
	entries, err := core.ScanMarkdownFiles(dirPath)
	if err != nil {
		s.files = []core.FileEntry{}
	} else {
		s.files = entries
	}
	s.selectedIdx = -1
	s.list.Refresh()
}

// selectedPath returns the currently selected file's path, or empty string.
func (s *sidebarCtrl) selectedPath() string {
	if s.selectedIdx >= 0 && s.selectedIdx < len(s.files) {
		return s.files[s.selectedIdx].Path
	}
	return ""
}

// toggle shows/hides the sidebar
func (s *sidebarCtrl) toggle() {
	s.hidden = !s.hidden
	if s.parent != nil {
		if s.hidden {
			s.parent.Hide()
		} else {
			s.parent.Show()
		}
	}
}
