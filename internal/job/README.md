# Job

定时任务模块

## 编译

```bash
go build -o bin/job/decert-job internal/job/cmd/main.go
```

## 运行环境配置

```bash
cp ./internal/job/cmd/config.demo.yaml ./bin/job/config.yaml
vi ./bin/job/config.yaml
```

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


### 区块链配置

配置项：
```yaml
# blockchain configuration
blockchain:
  sign-private-key: ""
  airdrop-private-key: ""
  chain-id: 5
  attempt: 200
  provider:
    - url: "https://rpc.ankr.com/eth_goerli"
      weight: 5
    - url: "https://rpc.ankr.com/eth_goerli"
      weight: 10
```

sign-private-key: 签名私钥
airdrop-private-key: 空投私钥
chain-id: 链ID
attempt: 尝试次数
provider：RPC
url：RPC地址
weight：权重

### 合约配置(V1)

配置项：
```yaml
contract:
  badge: ""
  quest: ""
  quest-minter: ""
```

badge: badge 合约地址
quest: quest 合约地址
quest-minter: quest minter 合约地址

### 合约配置(V2多链)

配置项：
```yaml
contract-v2:
  137:
    badge: "0xaD1789A4cA2640a570a1268c84EA87e09A10A1f5"
    quest: "0xd1E50047cEbaD1d826c3F7961A35048B518652F2"
    quest-minter: "0xeb4E2EE6f165e7E1B9a4F1c4532C0433a6Ef397B"
    provider:
      - url: "https://rpc.ankr.com/eth_goerli"
        weight: 5
      - url: "https://rpc.ankr.com/eth_goerli"
        weight: 10
```

137：链ID
badge: badge 合约地址
quest: quest 合约地址
quest-minter: quest minter 合约地址
provider：RPC

### IPFS 配置

配置项：
```yaml
# ipfs configuration
ipfs:
  api: "https://ipfs.io/ipfs/"
```

api：IPFS 节点地址