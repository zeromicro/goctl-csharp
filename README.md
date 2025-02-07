# goctl-csharp

![go-zero](https://img.shields.io/badge/Github-go--zero-brightgreen?link=https://github.com/zeromicro/go-zero&logo=github)
![License](https://img.shields.io/badge/License-MIT-blue?link=https://github.com/zeromicro/goctl-csharp/blob/main/LICENSE)

goctl-csharp 是一款基于 goctl 的插件，用于生成 C# （.Net 6 及以上版本） 使用的 http 客户端。

```bash
# 安装 goctl-csharp 插件
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/goctl-csharp
```

```bash
# 生成代码
goctl api plugin -plugin goctl-csharp="csharp -namespace Demo.NS" -api demo.api -dir .
```