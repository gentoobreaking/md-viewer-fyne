package core

import (
	"os"
	"path/filepath"
	"strings"
)

// FileEntry represents a file in the sidebar
type FileEntry struct {
	Name string
	Path string
}

// ReadFile reads a file and returns its content as string
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// BaseName returns the file name from a path
func BaseName(path string) string {
	return filepath.Base(path)
}

// ScanMarkdownFiles scans a directory for .md files
func ScanMarkdownFiles(dirPath string) ([]FileEntry, error) {
	// Clean the path — Fyne may return file:// URIs
	dirPath = strings.TrimPrefix(dirPath, "file://")

	var entries []FileEntry

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip hidden directories
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			entries = append(entries, FileEntry{
				Name: info.Name(),
				Path: path,
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return entries, nil
}
