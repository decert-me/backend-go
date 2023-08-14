FROM debian:latest

ENV DEBIAN_FRONTEND=noninteractive

# 设置Golang版本
ARG GOLANG_VERSION=1.20.7

# 下载Golang二进制文件并安装
RUN apt-get update && apt-get install -y wget && apt-get update && apt-get install -y python3
RUN wget -q https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O /tmp/golang.tar.gz && \
    tar -C /usr/local -xzf /tmp/golang.tar.gz && \
    rm /tmp/golang.tar.gz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

WORKDIR $GOPATH/src

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o executorserver ./cmd/executorserver

EXPOSE 5050/tcp 5051/tcp

ENTRYPOINT ["./executorserver"]
