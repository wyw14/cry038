# Bug 是什么

同一准备项连续发生状态变化时，操作时间线只保留最后一项；查询返回对象还能反向修改仓储中的既有审计记录。

# 如何触发

运行 `go test ./internal/... -run '^TestTimelineHistoryPipeline' -count=20`。

# 错误信息

测试会报告 `state history was overwritten` 或 `repository truncated audit history`，并可能报告读取结果污染已存时间线。
