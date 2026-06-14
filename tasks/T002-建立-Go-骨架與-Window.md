# T002 - 建立 Go 骨架與 Window

- **Status**: done
- **Type**: Development
- **Assignee**: 寶寶
- **Parent**: T001
- **Created**: 2026-04-23
- **Completed**: 2026-04-23

## 目標
在 md-viewer-app 專案建立基本的 Go 結構，包含主程式骨架與基本 Window

## 驗收條件
- [x] go.mod 建立完成，依賴 Fyne + goldmark
- [x] main.go 能編譯運行
- [x] Window 有基本的側邊欄和預覽區域佈局
- [x] 預覽區預設顯示歡迎頁

## 實際產出

### 目錄結構
```
md-viewer-app/
├── main.go              ← 入口
├── go.mod / go.sum      ← Go module
├── FyneApp.toml         ← Fyne 打包元數據
├── ui/
│   ├── window.go        ← 主視窗、選單列、快捷鍵
│   ├── sidebar.go       ← 側邊欄檔案列表
│   ├── preview.go       ← Markdown 預覽區
│   └── theme.go         ← 主題（深色模式）
├── core/
│   └── file_tree.go     ← 檔案讀取 + 資料夾掃描
└── assets/              ← （預留）
```

### 功能實作
| 功能 | 狀態 | 說明 |
|------|------|------|
| 主視窗 | ✅ | 900×600，HSplit 佈局 |
| 側邊欄 | ✅ | 檔案列表 + 點擊切換預覽 |
| 預覽區 | ✅ | `widget.NewRichTextFromMarkdown()` 內建 goldmark |
| 選單列 | ✅ | 檔案 / 顯示 / 說明 |
| 快捷鍵 | ✅ | ⌘O / ⌘⇧O / ⌘B |
| 深色模式 | ✅ | 自動跟隨 macOS（Fyne 內建） |
| .app 打包 | ✅ | 31MB .app，啟動正常 |

### API 修正記錄
- `widget.NewMarkdown()` → 不存在，正確是 `widget.NewRichTextFromMarkdown()`
- `dialog.ShowFolderOpenFunc` → 正確是 `dialog.ShowFolderOpen`
- `fyne.Shortcut{}` → 不存在，正確是 `desktop.CustomShortcut{}`
- Fyne 模組名 `md-viewer-app` 會變成 exe 名 → 用 FyneApp.toml 覆蓋

---
github_issue: https://github.com/openclawchen8-lgtm/openclaw-tasks/issues/95

*任務完成：寶寶 | 2026-04-23*