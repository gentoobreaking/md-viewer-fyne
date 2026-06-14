github_issue: https://github.com/openclawchen8-lgtm/openclaw-tasks/issues/94
# T006 - Build 與測試 macOS App

- **Status**: pending
- **Type**: Build
- **Assignee**: 寶寶
- **Parent**: T003, T004, T005
- **Created**: 2026-04-23

## 目標
將 md-viewer-app 打包成 macOS .app 並完成基本功能測試

## 驗收條件
- [ ] `fyne package` 成功產生 .app bundle
- [ ] .app 雙擊可啟動，Window 正常顯示
- [ ] Open File 功能正常（Open... / Cmd+O）
- [ ] Markdown 預覽正常渲染
- [ ] 側邊欄檔案列表正常運作
- [ ] 沒有 Crash 或 runtime panic

## Build 命令
```bash
fyne package -icon icon.png -category public.app-category.productivity
```

## 備註
- 需先確認 T001-T005 全部完成
- 若 fyne package 有問題，記錄錯誤並考慮替代打包方式（`go build` + macOS app bundle 結構）