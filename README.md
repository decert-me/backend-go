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
go build
```
## 运行
```bash
# 复制配置文件demo
cp ./config/config.demo.yaml ./config/config.yaml
# 修改配置文件
vi ./config/config.yaml
# 运行
./backend-go (windows运行命令为 backend-go.exe)
```
## 测试
```bash
go test ./internal/app/service
go test ./internal/app/service/blockchain/
go test ./pkg/...
```

