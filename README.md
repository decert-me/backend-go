# backend-go
![](https://img.shields.io/badge/license-MIT-green)
[![codecov](https://codecov.io/gh/decert-me/backend-go/branch/feature/testing/graph/badge.svg?token=D68XAECVLI)](https://codecov.io/gh/decert-me/backend-go)
[![goreportcard for backend-go](https://goreportcard.com/badge/github.com/decert-me/backend-go)](https://goreportcard.com/report/github.com/decert-me/backend-go)
## 安装
```bash
git clone https://github.com/decert-me/backend-go.git
```
## 编译
```bash
# 主程序
go build -o bin/app/decert-app internal/app/cmd/main.go

# 定时处理程序
go build -o bin/job/job-app internal/job/cmd/main.go
```
## 配置
```bash
# 主程序配置
cp ./internal/app/cmd/config.demo.yaml ./bin/app/config.yaml
vi ./bin/app/config.yaml

# 定时处理程序配置
cp ./internal/job/cmd/config.demo.yaml ./bin/job/config.yaml
vi ./bin/job/config.yaml
```
## 运行
```bash
# 主程序
cd bin/app
./decert-app

# 定时处理程序
cd bin/job
./job-app
```

## 测试
```bash
go test ./internal/app/service
go test ./internal/job/service
go test ./pkg/...
```

