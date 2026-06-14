# AGENTS.md - 專案協作規範

## 基本資料
- **專案：** md-viewer-app
- **目錄：** `/Users/claw/Tasks/md-viewer-app`
- **技術棧：** Go + Fyne + goldmark
- **目標：** macOS 原生 Markdown 預覽 App

---

## 開發原則

### 每次變更前
1. 確認 `T*.md` 任務狀態
2. `git pull`（若有多端協作）
3. 修改後執行 `go build` 確認編譯通過

### 任務完成後（強制）
1. 更新 T*.md 的 `status: done` + 勾選所有完成項目
2. 更新 README.md 對應任務列為 `✅ done`
3. `git add` + `git commit` 並標注任務編號

### Git 使用
- **分支策略：** 主分支开发，新建功能前开 feature branch
- **Commit 格式：** `[T001] 摘要`，例：`[T001] 安裝 Fyne CLI 並建立測試程式`
- **禁止：** 未經溝通直接 push 到 main

---

## 技術決策記錄

| 日期 | 決策 | 理由 |
|------|------|------|
| 2026-04-23 | 優先選 Fyne GUI 框架 | 純 Go、可 build .app、不需 Xcode GUI |
| 2026-04-23 | 使用 goldmark 解析 Markdown | Go 生態、支援 GFM、MIT License |

---

## 相關專案

| 專案 | 位置 | 關係 |
|------|------|------|
| md-viewer (CLI) | `/Users/claw/Projects/md-viewer/` | 技術參考（goldmark 使用方式） |

---

*最後更新：2026-04-23 by 寶寶*