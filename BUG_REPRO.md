# Bug 是什么

登记替代材料实际消耗时，原计划数量和替代用量会在领域、应用、汇总及导出链路中重复累加。

# 如何触发

运行 `go test ./internal/... -run '^TestReplacementAccounting' -count=20`。

# 错误信息

计划量为 20、替代用量为 8 时，测试会看到 `replacement consumption=28`、应用结果为 36，或导出/汇总返回 28 而不是 8。
