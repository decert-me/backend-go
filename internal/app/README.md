# App

## 安装

```bash
git clone https://github.com/decert-me/backend-go.git```]
```

## 编译

```bash
go build -o bin/app/decert-app internal/app/cmd/main.go
```

## 运行环境配置

```bash
cp ./internal/app/cmd/locale.json ./bin/app/locale.json
cp ./internal/app/cmd/config.demo.yaml ./bin/app/config.yaml
vi ./bin/app/config.yaml
```

### 运行端口配置

配置项：

```
# system configuration
system:
  env: develop
  addr: 8888
  i18n: "locale.json"
  website: ""
```

env：运行环境，可选值为 develop、test、production
addr：运行端口
i18n：国际化配置文件
website：前端网站index.html文件路径，用于代理meta标签

### 数据库配置

配置项：
```yaml
# pgsql configuration
pgsql:
  path: "127.0.0.1"
  port: "5432"
  config: ""
  db-name: ""
  username: "postgres"
  password: "123456"
  auto-migrate: true
  prefix: ""
  slow-threshold: 200
  max-idle-conns: 10
  max-open-conns: 100
  log-mode: "info"
  log-zap: false
```

path：数据库地址
port：数据库端口
config：数据库配置
db-name：数据库名称
username：数据库用户名
password：数据库密码
auto-migrate：是否自动迁移数据库
prefix：数据库表前缀
slow-threshold：慢查询阈值，单位毫秒
max-idle-conns：最大空闲连接数
max-open-conns：最大连接数
log-mode：日志级别
log-zap：是否使用zap日志库


### Redis 配置

配置项：
```yaml
# redis configuration
redis:
  db: 0
  addr: 127.0.0.1:6379
  password: ""
  prefix: "decert:"
```

db：Redis数据库
addr：Redis地址和端口
password：密码
prefix：前缀

### 日志级别配置

配置项：
```yaml
# log configuration
log:
  level: info
  save: true
  format: console
  log-in-console: true
  prefix: '[backend-go]'
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

配置项：

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

### 文件上传配置

配置项：

```yaml
# local configuration
local:
  path: 'uploads/file'
  ipfs: 'uploads/ipfs'
```

path：本地文件保存路径
ipfs：IPFS文件保存路径

### 区块链配置

配置项：

```yaml
# blockchain configuration
blockchain:
  ens-rpc: "https://rpc.ankr.com/eth"
  sign-private-key: ""
  signature: "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\n"
  signature-exp: 240
```

ens-rpc：ENS查找RPC地址
sign-private-key：签名私钥
signature：签名内容
signature-exp：签名过期时间，单位秒

### 合约配置

配置项：

```yaml
# contract configuration
contract:
  badge: "0xaD1789A4cA2640a570a1268c84EA87e09A10A1f5"
  quest: "0xd1E50047cEbaD1d826c3F7961A35048B518652F2"
  quest-minter: "0xeb4E2EE6f165e7E1B9a4F1c4532C0433a6Ef397B"
```

badge：badge合约地址
quest：quest合约地址
quest-minter：quest-minter合约地址

### 挑战信息配置

配置项：

```yaml
# quest configuration
quest:
  encrypt-key: "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca"
```

encrypt-key：挑战信息加密密钥

### IPFS 配置

配置项：

```yaml
# ipfs configuration
ipfs:
  - api: "https://ipfs.io/ipfs"
    upload-api: "http://192.168.1.10:3022/v1"
```

api：IPFS节点地址
upload-api：IPFS上传API地址

### 代码运行配置

配置项：

```yaml
# judge configuration
judge:
  judge-api:
    - url: "http://192.168.1.10:3022/v1"
      weight: 10
```

url：代码运行API地址
weight：权重

### 空投配置

配置项：

```yaml
share:
  verify-key: "123456"                           # 校验key
  callback: "http://192.168.1.10:8105"   # 回调接口
```

verify-key：校验key
callback：回调接口

### discord 消息通知配置

配置项：

```yaml
# discord configuration
discord:
  active: true
  token: "TzATY3NA.G3oGb.VryUgxsBM"
  success-channel: "1105777577794211840"
  failed-channel: "1105777577794211840"
```

active：是否启用
token：机器人token
success-channel：成功通知频道ID
failed-channel：失败通知频道ID

### sentry 配置

配置项：

```yaml
# sentry configuration
sentry:
  dsn: "https://xxxx@xxx.ingest.sentry.io/xxxx"
  traces-sample-rate: 1.0       # Percentage of transactions to be traced.
  enable-tracing: true          # Enable performance tracing.
```

dsn：sentry dsn
traces-sample-rate：采样率
enable-tracing：是否启用性能追踪

### 社交账号绑定配置

配置项：

```yaml
social:
  wechat:
    api-key: "test"
    call-url: "http://127.0.0.1:8110"
  discord:
    api-key: "test"
    call-url: "http://127.0.0.1:8110"
```

wechat：微信配置
discord：discord配置

api-key：API密钥，与Auth配置一致
call-url：回调地址，Auth接口地址