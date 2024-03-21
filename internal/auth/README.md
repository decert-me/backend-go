# Auth

授权登录模块

## 编译

```bash
go build -o bin/auth/decert-auth internal/auth/cmd/main.go
```

## 运行环境配置

```bash
cp ./internal/auth/cmd/config.demo.yaml ./bin/auth/config.yaml
vi ./bin/auth/config.yaml
```

### 运行端口配置

配置项：
```
# system configuration
system:
  env: develop
  addr: 8889
```
设置 addr 为程序运行端口

### 项目配置

配置微信和 Discord 绑定回调 URL和校验 KEY，支持配置多个项目

配置项：
```yaml
# project configuration
project:
  decert:
    api-key: "test"
    callback-url: "http://127.0.0.1:8087/api"
```
1、api-key 需与decert app 模块 social configuration 配置项保持一致

2、callback-url 填写 decert app 模块接口 URL


### 微信公众号配置

微信公众号用户绑定配置

配置项：
```yaml
# auth configuration
auth:
  wechat:
    app-id: ""                    # wechat app id
    app-secret: ""                # wechat app secret
    token: ""                     # wechat token
    encoding-aes-key: ""          # wechat encoding aes key(可选)
```

进入 [微信公众号基本配置](https://mp.weixin.qq.com/advanced/advanced?action=dev&t=advanced/dev&token=678258990&lang=zh_CN)

获取 AppID 和 AppSecret 配置：
![](https://ipfs.io/ipfs/bafybeihmyjqidtjtjt5yulxuxclfzvaucgg736fkana4wqsrntrl4w47pq/Snipaste_2023-12-19_16-43-37.png)

服务器配置：

1、部署 Auth 程序，程序的 Token 和 EncodingAESKey 与微信公众号配置保持一致

2、到微信公众号服务器配置中，配置 URL、Token 和 EncodingAESKey

(https://ipfs.io/ipfs/QmQFgHGLDBUqR8ezUpxCutKwe6mvZF7EA2G3Ln8MmLhcaX/Snipaste_2023-12-19_16-39-51.png)

### Discord 配置

Discord 用户绑定配置

配置项：
```yaml
# auth configuration
auth:
  discord:
    client-id: ""
    client-secret: ""
```

1、进入 [DEVELOPER PORTAL](https://discord.com/developers/applications)

2、选择用户需要授权绑定的 Applications，未创建可以点击 New Application

3、进入 Applications -> Oauth2 页面复制 CLIENT ID 和 CLIENT SECRET 写入配置文件

4、Redirects 添加 decert 回调链接 https://decert.me/callback/discord

![](https://ipfs.io/ipfs/bafybeihw43kchnyicpn2m7y4o2eb6ma4sgkksaymcx6osm7igcscc35wpe/Snipaste_2023-12-19_17-14-08.png)

### Github 配置

Github 用户绑定配置

配置项：
```yaml
# auth configuration
auth:
  github:
    client-id: ""
    client-secret: ""
```

1、进入 [Github OAuth Apps](https://github.com/settings/developers)

2、选择用户需要授权绑定的 Apps，未创建可以点击 New OAuth App

3、进入 Apps 页面复制 CLIENT ID 和 CLIENT SECRET 写入配置文件

4、Authorization callback URL 添加 decert 回调链接 https://decert.me/