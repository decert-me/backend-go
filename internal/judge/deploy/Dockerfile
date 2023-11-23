FROM debian:latest

#RUN sed -E -i -e 's/(archive|ports).ubuntu.com/mirrors.aliyun.com/g' -e '/security.ubuntu.com/d' /etc/apt/sources.list
ENV DEBIAN_FRONTEND=noninteractive
#
#RUN apt-get update && apt-get install -y golang
#
#


# 设置Golang版本
ARG GOLANG_VERSION=1.20.7

# 下载Golang二进制文件并安装
RUN apt-get update && apt-get install -y wget
RUN wget -q https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz -O /tmp/golang.tar.gz && \
    tar -C /usr/local -xzf /tmp/golang.tar.gz && \
    rm /tmp/golang.tar.gz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:/usr/local/go/bin:$GOPATH/bin

WORKDIR $GOPATH/src


#RUN buildDeps='software-properties-common git libtool cmake python-dev python3-pip python-pip libseccomp-dev curl' && \
#    apt-get update && apt-get install -y python3 python-pkg-resources python3-pkg-resources $buildDeps && \
#    add-apt-repository ppa:openjdk-r/ppa && add-apt-repository ppa:longsleep/golang-backports && \
#    add-apt-repository ppa:ubuntu-toolchain-r/test && \
#    curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
#    apt-get update && apt-get install -y golang-go nodejs && \
#    pip3 install -i https://mirrors.aliyun.com/pypi/simple/ -I --no-cache-dir psutil gunicorn requests idna && \
#    apt-get purge -y --auto-remove $buildDeps && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o executorserver ./cmd/executorserver

EXPOSE 5050/tcp 5051/tcp

ENTRYPOINT ["./executorserver"]
