# 课堂准备与物料核验协同平台

平台把课程、课次、准备模板、负责人执行、风险锁定、耗材回填和课后复盘串成闭环，不包含考试或题库。模板生成课次时复制快照，后续模板改动不会污染历史；课堂开始后关键项禁止删除；存在未完成或阻塞的关键项时不能锁定开课确认。

配置见 `.env.example`。本地执行 PostgreSQL 迁移和 `scripts/seed.sql` 后运行 `go run ./cmd/server`，或执行 `docker compose up --build`。演示课次 `demo-print` 位于 Art-203。HTTP API 统一位于 `/api/v1`，例如 `POST /api/v1/sessions/demo-print/lock`；错误包含稳定 code、message 与 request_id。

目录中 `internal/domain` 保管课次快照和状态规则，`internal/application` 编排锁定与复盘建议，`internal/repository` 提供线程安全内存实现和 pgx 连接入口，`internal/platform` 生成离线审计导出。Vue 页面覆盖课次详情、清单、风险、模板、复盘、待办与课程总览。

实际验证命令：`go test ./...`、`go test -race ./...`、`go vet ./...`，以及 `cd web && npm ci && npm run test && npm run build`。健康与就绪探针为 `/healthz`、`/readyz`。
