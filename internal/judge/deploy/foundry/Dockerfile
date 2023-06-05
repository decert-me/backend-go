FROM ghcr.io/foundry-rs/foundry:latest

# 环境配置
RUN apk update && apk add git bash && git config --global user.email "you@example.com" && git config --global user.name "Your Name" && \
    mkdir /foundry && cd /foundry && forge init && forge test && \
    rm -r script src test

