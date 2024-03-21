# Judge

Solidity 代码运行和判题服务

## 配置项

### 运行端口配置

配置项：

```
# system configuration
system:
  env: develop
  addr: 8888
```

env：运行环境，可选值为 develop、public

addr：运行端口

### 日志级别配置

配置项：
```yaml
# log configuration
log:
  level: info
  save: true
  format: console
  log-in-console: true
  prefix: '[decert-judge]'
  director: log
  show-line: true
  encode-level: LowercaseColorLevelEncoder
  stacktrace-key: stacktrace
```

level：日志级别 debug、info、warn、error、dpanic、panic、fatal

save：是否保存日志

format：日志格式

log-in-console：是否在控制台输出日志

prefix：日志前缀

director：日志保存路径

show-line：是否显示行号

encode-level：日志编码级别

stacktrace-key：堆栈信息

### JWT 配置

配置项（需要与app程序配置保持一致）：

```yaml
# auth configuration
auth:
  signing-key: "Decert"
  expires-time: 86400
  issuer: "Decert"
```

signing-key：签名密钥

expires-time：过期时间，单位秒

issuer：签发人

### 挑战信息配置

配置项：

```yaml
# quest configuration
quest:
  encrypt-key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca"
```

encrypt-key：挑战信息加密密钥

### Docker 配置

配置项：

```yaml
# docker configuration
docker:
  clear-enabled: true 
  clear-time: 15
```

clear-enabled：是否开启 Docker 定时清空容器

clear-time: Docker 定时清空超过闲置时长的容器（分钟）

### 代码运行配置（未启用）

配置项：

```yaml
judge:
  sandbox-service: "http://192.168.1.15:5050/"  # 沙盒服务URL
  work-path: "/Users/mac/Code/resource/"        # 判题模块工作目录，临时保存用户代码
  cache-path: "/Users/mac/Code/resource/cache"  # 判题模块缓存目录，缓存 solc 等
  javascript-path: "/Users/mac/.nvm/versions/node/v18.16.0/bin/node"
  typescript-path: "/usr/local/bin/ts-node"
  golang-path: "/usr/local/go/bin/go"
  python-path: "/usr/bin/python3"
```

sandbox-service：沙盒服务URL，请查看 sandbox 项目

work-path：判题模块工作目录，临时保存用户代码

cache-path：判题模块缓存目录，缓存 solc 等

javascript-path：nodejs 路径

typescript-path：ts-node 路径

golang-path：golang 路径

python-path：python 路径