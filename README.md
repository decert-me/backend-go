# backend-go
![](https://img.shields.io/badge/license-MIT-green)
[![codecov](https://codecov.io/gh/decert-me/backend-go/branch/feature/testing/graph/badge.svg?token=D68XAECVLI)](https://codecov.io/gh/decert-me/backend-go)
[![goreportcard for backend-go](https://goreportcard.com/badge/github.com/decert-me/backend-go)](https://goreportcard.com/report/github.com/decert-me/backend-go)
## 安装
```bash
git clone https://github.com/decert-me/backend-go.git
```
## 运行环境

```shell
- Golang >= v1.19 && < v1.21
- Redis
- PostgreSQL
- Docker
```

## 环境配置
1、安装 Docker
脚本一键安装: `sudo curl -sSL https://get.daocloud.io/docker | sh`

详细步骤参照： https://docs.docker.com/install/
## 编译

```bash
# 主程序
go build -o bin/app/decert-app internal/app/cmd/main.go

# 定时任务程序
go build -o bin/job/decert-job internal/job/cmd/main.go

# 授权登录程序
go build -o bin/auth/decert-auth internal/auth/cmd/main.go

# 判题程序
go build -o bin/judge/decert-judge internal/judge/cmd/main.go
```

## 配置

查看各个项目配置说明：
- [主程序配置说明](./internal/app/README.md)
- [定时任务程序配置说明](./internal/job/README.md)
- [授权登录程序配置说明](./internal/auth/README.md)
- [判题程序配置说明](./internal/judge/README.md)


```bash
# 主程序配置
cp ./internal/app/cmd/config.demo.yaml ./bin/app/config.yaml
vi ./bin/app/config.yaml

# 定时任务程序配置
cp ./internal/job/cmd/config.demo.yaml ./bin/job/config.yaml
vi ./bin/job/config.yaml

# 授权登录程序
cp ./internal/auth/cmd/config.demo.yaml ./bin/auth/config.yaml
vi ./bin/auth/config.yaml

# 判题程序配置
cp ./internal/judge/cmd/config.demo.yaml ./bin/judge/config.yaml
vi ./bin/judge/config.yaml
```

## Docker 构建判题镜像

```shell
sudo docker build -t judge:1.1 -f internal/judge/Dockerfile .
```

## 运行

```bash
# 主程序
cd bin/app
./decert-app

# 定时任务程序
cd bin/job
./decert-job

# 授权登录程序
cd bin/auth
./decert-auth

# 判题程序
cd bin/judge
./decert-judge
```