# md-viewer-app

**純 Markdown 預覽 macOS App，用 Go 開發。**

> 動機：MacDown 無法預設只開預覽，需要一個乾淨的 Markdown 閱讀工具。

---

## 🎯 專案目標

建立一個 macOS 原生 App，功能：
- 📂 拖放 / 開啟 .md 檔，純預覽（零編輯區）
- 🗂️ 側邊欄檔案列表
- 🌙 深色模式（跟隨系統）
- 🔖 書籤 / 捷徑
- 📁 支援拖放資料夾，掃描所有 .md

---

## ⚙️ 可行方案分析

以下為 2026-04-23 評估結果：

| 方案 | 技術棧 | 複雜度 | UI 品質 | 優點 | 缺點 |
|------|--------|--------|---------|------|------|
| **Fyne** | Go GUI 框架 + WebView | ★★☆ | ★★★ | 純 Go，不需要 Node.js/Xcode GUI，可直接 build 成 `.app` | WebView 渲染取決於系統版本，部分 CSS 可能不支援 |
| **Wails** | Go + WebView（Electron 底層） | ★★☆ | ★★★★ | 功能完整，支援系統通知、系統 tray | 需要另外安裝 npm/wails CLI，學習曲線稍高 |
| **Cocoa + Go** | Go + Objective-C 綁定 | ★★★★ | ★★★★★ | 原生 Cocoa 體驗，最佳 UI | 需要 Objective-C bridging，CocoaPods 環境複雜，Xcode GUI 依賴高 |
| **go-astilectron** | Go + Chromium (Electron) | ★★★ | ★★★★ | 豐富的前端生態，UI 靈活 | 依賴 Chromium binary，App bundle 體積大（>100MB） |
| **HTML/CLI 預覽器** | Python/Go CLI + 瀏覽器 | ★☆☆ | ★★★ | 最快實作，無需 GUI | 不是原生 App，需要終端操作 |

### 🏆 最終判斷：優先選 **Fyne**，其次 **Wails**

**理由（Fyne）：**
1. 純 Go，从代码编写到最终构建都是同一种语言，开发体验一致
2. 包含 WebView 支持，可以用现代 HTML/CSS 做渲染，UI 好看
3. 可直接 build 成 macOS `.app`，不需要 Xcode GUI 打包
4. 已有成功案例（Fyne 本身是知名开源项目）
5. Go 生态中有 goldmark 做 Markdown 解析，质量高

**风险点：**
- Fyne WebView 性能取决于 macOS 系统 WebKit 版本
- 需要确认 `fyne bundle` / `fyne package` 能生成可用 .app

**次選（Wails）：**
- 如果 Fyne WebView 体验不佳，可以切到 Wails
- Wails 用真实 Chromium，UI 一致性更好

---

## 📦 技術棧（目前規劃）

- **語言**：Go 1.21+
- **Markdown 解析**：github.com/yuin/goldmark
- **GUI 框架**：Fyne 2.x（含 WebView）
- **構建工具**：go build + fyne 命令
- **目標平台**：macOS（.app bundle）

---

## 📁 專案結構

```
md-viewer-app/
├── README.md          ← 你現在看的
├── SPEC.md            ← 詳細需求規格（待建立）
├── AGENTS.md          ← 團隊協作規範
└── tasks/             ← 任務追蹤
    ├── T001-評估-Fyne-環境.md
    ├── T002-建立-Go-骨架與-Window.md
    ├── T003-實作-Markdown-解析與-WebView-渲染.md
    ├── T004-實作-側邊欄-檔案列表.md
    ├── T005-實作-深色模式支援.md
    └── T006-Build-與-測試-macOS-app.md
```

---

## 🔗 關聯專案

| 專案 | 關係 |
|------|------|
| `/Users/claw/Projects/md-viewer` | CLI 版本的 Markdown 預覽器（Go + goldmark），本專案的技術參考 |

---

*建立時間：2026-04-23 by 寶寶*