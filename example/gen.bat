@echo off

for /f "delims=" %%t in ('pwd') do set here_dir=%%t

cd %here_dir%/server

@rem 生成 Go 服务端代码
go mod tidy
goctl api go --api demo.api --dir .

@rem 生成 C# 客户端代码
goctl api plugin -plugin goctl-csharp="csharp -namespace GenLib" -api demo.api -dir ../client/GenLib

cd %here_dir%