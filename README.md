# bili-grpc-api-go
[XiaoMiku01/bilibili-grpc-api-go](https://github.com/XiaoMiku01/bilibili-grpc-api-go) 长时间没有维护了 本项目与其类似

接口来自于[SocialSisterYi/bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect)

## 安装

```bash
go get -u github.com/stmtc233/bili-grpc-api-go
```

## 切换版本

大陆版、国际版和 TV 版使用相同的 Go module 路径和包导入路径。调用代码不需要改动，
只需要在 `go get` 时选择对应的分支或 Release；切换后运行 `go mod tidy`。

```bash
# 大陆版
go get github.com/stmtc233/bili-grpc-api-go@main

# 国际版
go get github.com/stmtc233/bili-grpc-api-go@international

# TV 版
go get github.com/stmtc233/bili-grpc-api-go@tv
```

也可以固定到已发布版本：

```bash
go get github.com/stmtc233/bili-grpc-api-go@v1.0.12
go get github.com/stmtc233/bili-grpc-api-go@international-v1.0.1
go get github.com/stmtc233/bili-grpc-api-go@tv-v1.0.1
go mod tidy
```

例如 `bilibili/app/view/v1` 在三个版本中的导入路径都不变；Go module 依赖版本决定实际使用的
接口定义。一个 Go module 不能同时引入这三个同路径版本，需要同时对比时请使用不同工作目录。

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

`Sync Upstream Proto And Release` 每天自动检查
[`stmtc233/bapis-proto`](https://github.com/stmtc233/bapis-proto) 的三个分支：

- `main`：大陆版
- `international`：国际版
- `tv`：TV 版

每个版本使用独立分支。首次运行会自动创建缺失的分支；上游有变化时，workflow
同步 `.proto`、重新生成 `.pb.go` 和 `_grpc.pb.go`、运行 `go test ./...`，然后直接提交、推送并创建 GitHub Release。
没有变化时不会产生空提交或空 Release。Release tag 使用 `v*`、`international-v*`
和 `tv-v*` 前缀，避免三个版本互相覆盖。

workflow 也支持手动选择单个版本进行补发，但正常更新不需要人工合并 PR。

### 注意

- 如果仓库开启了严格的分支保护，需要允许 GitHub Actions 推送生成后的提交和 tag。
- 自动发布 tag 默认按现有最新 tag 的补丁版本递增，例如 `v1.0.4 -> v1.0.5`。
- `.proto-sync-manifest` 只记录当前版本从上游同步过的文件，后续同步会按这个清理上游已删除的 proto。
