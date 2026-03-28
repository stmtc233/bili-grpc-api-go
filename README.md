# bili-grpc-api-go
[XiaoMiku01/bilibili-grpc-api-go](https://github.com/XiaoMiku01/bilibili-grpc-api-go) 长时间没有维护了 本项目与其类似

接口来自于[SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)

## 安装

```bash
go get -u github.com/stmtc233/bili-grpc-api-go
```

## 可视化 gRPC 调试工具

仓库内置了一个本地 Web 调试页，可以用来快速选择接口、填写参数并查看 protobuf 解析后的返回值。

### 启动

```bash
go run ./cmd/grpc-debugger
```

默认会监听在 `http://127.0.0.1:8090`。

也可以自定义监听地址：

```bash
go run ./cmd/grpc-debugger -listen :9000
```

### 功能

- 左侧搜索并选择仓库里已注册的 gRPC service / method
- 自动读取 protobuf 描述并生成请求表单
- 支持切换到 JSON 模式，直接手写复杂请求体
- 支持填写目标地址、TLS、authority、server name、metadata、超时
- 返回值会以格式化 JSON 展示，同时显示规范化后的请求体
- 流式接口会被标记出来，但当前版本只支持一元 RPC 调试

### 使用步骤

1. 启动 `grpc-debugger`
2. 打开浏览器访问本地页面
3. 在左侧选中要调试的接口
4. 填写目标 gRPC 地址和必要的 metadata
5. 通过表单模式或 JSON 模式填写请求参数
6. 点击“调用接口”查看返回结果

### Metadata 示例

```json
{
  "authorization": "identify_v1 your-token",
  "x-bili-device-bin": "xxxxx",
  "x-bili-metadata-bin": "yyyyy"
}
```

## GitHub Actions

仓库现在可以把“同步 proto”和“生成发布”拆开处理：

1. `Sync Upstream Proto`
   从 `mikufans-project/bilibili-API-collect` 的 `grpc_api` 目录同步 `.proto` 文件，并自动创建一个同步 PR。
   默认每周一运行一次，也可以手动触发并指定上游分支或 commit。

2. `Generate Go Package And Release`
   基于当前仓库里的 `.proto` 重新生成 `.pb.go` / `_grpc.pb.go`，然后自动创建新 tag 并发布 GitHub Release。
   这个 workflow 会在 `main` 分支的 `.proto` 发生变更后自动运行，也支持手动指定 tag。
   如果你要基于同步分支或自定义分支发布，可以在手动触发时填写 `source_ref`。

### 建议流程

1. 先运行或等待 `Sync Upstream Proto` 创建同步 PR。
2. 在 PR 里按需保留你的自定义改动。
3. 合并到 `main` 后，`Generate Go Package And Release` 会自动生成并发布。
4. 如果还没合并，也可以手动运行 `Generate Go Package And Release`，并把 `source_ref` 指到同步分支或你的自定义分支。

### 注意

- 如果仓库开启了严格的分支保护，需要允许 GitHub Actions 推送生成后的提交和 tag。
- 自动发布 tag 默认按现有最新 tag 的补丁版本递增，例如 `v1.0.4 -> v1.0.5`。
- `.proto-sync-manifest` 只记录上游同步过的文件，后续同步会按这个清理上游已删除的 proto，尽量不误删你的自定义 proto。
