FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk update --no-cache && apk add --no-cache tzdata && apk add wget

WORKDIR /backend-go

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags="-s -w" -o /bin/judge/decert-judge internal/judge/cmd/main.go

FROM ubuntu:latest

WORKDIR /judge
COPY --from=builder /bin/judge/decert-judge /judge/decert-judge
COPY /bin/judge/config.yaml /judge/config.yaml
ADD solc_down.sh /judge/

# 环境配置
RUN sed -E -i -e 's/(archive|ports).ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' -e '/security.ubuntu.com/d' /etc/apt/sources.list && \
    buildDeps='curl' && apt-get update && apt-get install -y git $buildDeps && \
    curl -L https://foundry.paradigm.xyz | bash && . /root/.bashrc && foundryup && \
    git config --global user.email "you@example.com" && git config --global user.name "Your Name" && \
    mkdir -p /cloud-run-foundry/ && cd /cloud-run-foundry/ && forge init && \
    mkdir -p "$HOME/.svm" && cd /judge && chmod +x ./solc_down.sh && bash ./solc_down.sh && \
    apt-get purge -y --auto-remove $buildDeps && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# 运行应用程序
ADD entrypoint.sh /judge/

EXPOSE 8888
ENTRYPOINT /judge/entrypoint.sh