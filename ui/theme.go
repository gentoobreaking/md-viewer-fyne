package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// mdViewerTheme wraps the default Fyne theme with customizations
type mdViewerTheme struct {
	fyne.Theme
}

func (m *mdViewerTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Use default Fyne colors (auto dark/light mode)
	return theme.DefaultTheme().Color(name, variant)
}

func (m *mdViewerTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *mdViewerTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *mdViewerTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
