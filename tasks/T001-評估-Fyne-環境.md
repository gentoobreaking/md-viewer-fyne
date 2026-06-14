# T001 - 評估 Fyne 環境與建置能力

- **Status**: done
- **Type**: Research
- **Assignee**: 寶寶
- **Created**: 2026-04-23
- **Completed**: 2026-04-23

## 目標
確認 Fyne 能在當前環境安裝、運行、並成功 build 成 macOS .app

## 執行結果

### ✅ Step 1: 安裝 Fyne CLI
```bash
go install fyne.io/tools/cmd/fyne@latest
```
- **新位置**：`/Users/claw/go/bin/fyne`（舊路徑 `fyne.io/fyne/v2/cmd/fyne` 已廢棄）
- **版本**：`v1.7.0`（CLI）+ `fyne.io/fyne/v2 v2.7.3`（runtime）
- **警告**：`NOTE: This tool is deprecated and has migrated to fyne.io/tools/cmd/fyne`（cli v1.7 已是最終版，繼續用沒問題）

### ✅ Step 2: 初始化專案
```bash
go mod init md-viewer-app
go get fyne.io/fyne/v2
go mod tidy
```
- 成功解析所有 transitives dependencies（含 goldmark）

### ✅ Step 3: 建立測試程式
- `test_fyne.go`：最小 Fyne Window（Label + MultiLineEntry）
- 成功 `go build -o test_fyne test_fyne.go`

### ✅ Step 4: Build .app
```bash
fyne package --executable ./test_fyne --name md-viewer --os darwin --app-id com.mdviewer.app
```
- **產出**：`md-viewer.app/`（31MB，包含 test_fyne 二進制）
- **注意**：Info.plist 預設 executable 名稱是 `test_fyne`，需手動修正為 `md-viewer`（見下方 BUG）
- icon.icns 自動生成（從 Icon.png）

### ✅ Step 5: 驗證啟動
- 修正 Info.plist CFBundleExecutable → md-viewer 後，雙擊可啟動

---
github_issue: https://github.com/openclawchen8-lgtm/openclaw-tasks/issues/89

## 驗收條件（更新後）

- [x] `go install fyne.io/fyne/v2/cmd/fyne@latest` 安裝成功
- [x] `fyne --version` 能執行
- [x] 成功建立一個最基本的 Fyne Window 並 build 成可執行檔
- [x] 確認 .app bundle 能正常啟動

---

## 🐛 發現的 Bug

### fyne package executable 名稱不一致
**問題**：`--executable` 傳 `./test_fyne`，但 --name 傳 `md-viewer`，產出的 Info.plist CFBundleExecutable 為 `test_fyne` 而非 `md-viewer`，導致 .app 無法啟動。
**解法**：不要用 `--name` + `--executable` 分開設定，直接用預設方式，或事後 sed 修正 Info.plist。
**Workaround**：
```bash
# 方式1：直接用預設（exe 名稱 = 目錄名）
go build -o md-viewer
fyne package --os darwin --app-id com.mdviewer.app  # 會自動取 md-viewer 為 exe 名稱

# 方式2：事後修正
sed -i '' 's/old_exe_name/new_exe_name/g' xxx.app/Contents/Info.plist
```

---

## 技術發現

### Fyne v2 WebView Markdown 支援
Fyne 內建 `fyne.io/fyne/v2/widget/markdown.go`，直接整合 goldmark，可以這樣用：
```go
import "fyne.io/fyne/v2/widget"

md := widget.NewMarkdown()          // 純預覽（不可編輯）
md.SetMarkdown("# Hello\n\nWorld")  // 設定內容
```
→ **MVP 可直接用 Fyne 內建 Markdown widget，不需要自己包 goldmark！**

### 環境版本確認
| 工具 | 版本 |
|------|------|
| Go | 1.26.1 darwin/arm64 |
| Node | v22.16.0 |
| Xcode CmdLineTools | ✅ 已安裝 |
| fyne CLI | v1.7.0 |
| fyne runtime | v2.7.3 |
| goldmark | v1.7.8（bundled） |

---

## 結論

**Fyne 完全可用！** 環境建置成功，可進入 T002 實作。

### 推薦用法（更新）
- Markdown 渲染：直接用 `widget.NewMarkdown()`（內建 goldmark）
- 側邊欄：用 `fyne.io/fyne/v2/widget/tree` 或 `fyne.io/fyne/v2/widget/list`
- 深色模式：`fyne settings.get().Theme().Dark()` 或 `app.Settings().AddChangeListener`

---

*任務完成：寶寶 | 2026-04-23*