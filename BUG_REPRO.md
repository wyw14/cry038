# Bug 是什么

关键安全检查处于阻塞状态时不会进入课前风险，并且课次仍可被锁定，HTTP 接口错误返回成功状态。

# 如何触发

运行 `go test ./internal/... -run '^TestBlockedCriticalRisk' -count=20`。

# 错误信息

测试会报告 `blocked critical item disappeared from risks`、`blocked risk missing from reminder` 或 `lock status=200`。
