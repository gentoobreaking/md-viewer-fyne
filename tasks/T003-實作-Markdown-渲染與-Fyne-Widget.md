---
github_issue: https://github.com/openclawchen8-lgtm/openclaw-tasks/issues/91
title: 實作 Markdown 渲染與 Fyne Widget
status: pending
assignee: 碼農1號
created: 2026-04-23
updated: 2026-04-24
depends: [T002]
---

## ⚠️ 注意：此任務已遷移至 md-viewer-webview

本專案（md-viewer-fyne）已拆分為獨立技術棧。此 T003 的 WebView 規劃已移到 `md-viewer-webview` 專案。

md-viewer-fyne 的 T003 將專注於 **Fyne 原生 Markdown 渲染**。

---

## 新目標：使用 Fyne 內建元件實作 Markdown 顯示

### Fyne NewRichTextFromMarkdown

Fyne 提供 `widget.NewRichTextFromMarkdown()` 方法，可以將 Markdown 轉為 RichText widget 顯示。

**限制**：
- 不支援表格
- 不支援 GFM 任務清單
- 不支援刪除線
- CSS 樣式受限於 Fyne 主題

### 子任務

### T003-A — 評估 Fyne Markdown 渲染極限
- **負責人**：碼農1號
- **內容**：
  - 測試 `widget.NewRichTextFromMarkdown()` 支援的功能
  - 列出已知限制
  - 評估是否需要自訂擴展
- **產出**：評估報告

### T003-B — 實作基本 Markdown 顯示
- **負責人**：碼農1號
- **內容**：
  - 在現有 `ui/preview.go` 中整合 Markdown 渲染
  - 支援：標題、粗體、斜體、程式碼、連結
- **產出**：更新 `ui/preview.go`

### T003-C — Fyne 主題與深色模式
- **負責人**：碼農1號
- **內容**：
  - 自訂 `fyne.Theme` 實作
  - 跟隨系統深色模式
- **產出**：更新 `ui/theme.go`

### T003-D — Build 與驗證
- **負責人**：碼農1號
- **內容**：
  - `go build -o md-viewer .`
  - `fyne package --os darwin --app-id com.mdviewer.fyne`
  - 測試 Markdown 顯示效果
- **產出**：.app bundle

---

## 備註

- **md-viewer-fyne vs md-viewer-webview**：詳見 `/Users/claw/Projects/COMPARISON.md`
- 此專案專注於 Fyne 原生 UI，犧牲 Markdown 渲染品質換取快速開發
- 如果需要精美 Markdown 顯示，建議使用 md-viewer-webview

---

*規劃：寶寶 | 2026-04-24*
