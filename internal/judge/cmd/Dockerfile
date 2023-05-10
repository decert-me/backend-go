FROM ubuntu:18.04

COPY build/java_policy /etc
RUN buildDeps='software-properties-common git libtool cmake python-dev python3-pip python-pip libseccomp-dev curl' && \
    apt-get update && apt-get install -y golang-go openjdk-11-jdk php-cli nodejs gcc-9 g++-9 && \
    update-alternatives --install  /usr/bin/gcc gcc /usr/bin/gcc-9 40 && \
    update-alternatives --install  /usr/bin/g++ g++ /usr/bin/g++-9 40 && \
    phpJitOption='opcache.enable=1\nopcache.enable_cli=1\nopcache.jit=1205\nopcache.jit_buffer_size=64M' && \
    echo $phpJitOption > /etc/php/8.0/cli/conf.d/10-opcache-jit.ini && \
    pip3 install -i https://mirrors.aliyun.com/pypi/simple/ -I --no-cache-dir psutil gunicorn flask requests idna && \
    cd /tmp && git clone -b newnew  --depth 1 https://gitee.com/qduoj/Judger.git && cd Judger && \
    mkdir build && cd build && cmake .. && make && make install && cd ../bindings/Python && python3 setup.py install && \
    apt-get purge -y --auto-remove $buildDeps && \
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    mkdir -p /code && \
    useradd -u 12001 compiler && useradd -u 12002 code && useradd -u 12003 spj && usermod -a -G code spj