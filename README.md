# bili-grpc-api-go
[XiaoMiku01/bilibili-grpc-api-go](https://github.com/XiaoMiku01/bilibili-grpc-api-go) 长时间没有维护了 本项目与其类似

接口来自于[SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)

## 安装

```bash
go get -u github.com/stmtc233/bili-grpc-api-go
```

## GitHub Actions

仓库现在可以把“同步 proto”和“生成发布”拆开处理：

1. `Sync Upstream Proto`
   从 `mikufans-project/bilibili-API-collect` 的 `grpc_api` 目录同步 `.proto` 文件，并自动创建一个同步 PR。
   默认每周一运行一次，也可以手动触发并指定上游分支或 commit。

2. `Generate Go Package And Release`
   基于当前仓库里的 `.proto` 重新生成 `.pb.go` / `_grpc.pb.go`，然后自动创建新 tag 并发布 GitHub Release。
   这个 workflow 会在 `main` 分支的 `.proto` 发生变更后自动运行，也支持手动指定 tag。

### 建议流程

1. 先运行或等待 `Sync Upstream Proto` 创建同步 PR。
2. 在 PR 里按需保留你的自定义改动。
3. 合并到 `main` 后，`Generate Go Package And Release` 会自动生成并发布。

### 注意

- 如果仓库开启了严格的分支保护，需要允许 GitHub Actions 推送生成后的提交和 tag。
- 自动发布 tag 默认按现有最新 tag 的补丁版本递增，例如 `v1.0.4 -> v1.0.5`。
- `.proto-sync-manifest` 只记录上游同步过的文件，后续同步会按这个清理上游已删除的 proto，尽量不误删你的自定义 proto。
