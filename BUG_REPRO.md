# Bug 是什么

从可复用准备模板生成并保存课次后，条目的模板键、默认负责人、关键标记和数量会丢失，后续提醒与导出也显示错误结果。

# 如何触发

运行 `go test ./internal/... -run '^TestTemplateSnapshotPipeline' -count=20`。

# 错误信息

测试会报告 `template defaults were not snapshotted`、`stored snapshot metadata was lost`、`critical template risk missing from reminder` 和 `export lost snapshotted owner`。
