# Bug 是什么

模板改进建议和课后复盘缺少课次隔离：不同课次互相判重，复盘内容串扰或被覆盖，导出只保留最后一项。

# 如何触发

运行 `go test ./internal/... -run '^TestReviewIsolation' -count=20`。

# 错误信息

测试会报告 `another session was treated as duplicate`、`review history was overwritten`、`reviews crossed session boundary` 或 `export omitted review history`。
