FROM ghcr.io/foundry-rs/foundry:latest

# 环境配置
RUN apk update && apk add --no-cache git bash

# 配置git
RUN git config --global --add safe.directory "*" && \
    git config --global user.email "you@example.com" && \
    git config --global user.name "Your Name"

# 设置foundry项目
WORKDIR /foundry
RUN forge init && \
    forge install OpenZeppelin/openzeppelin-contracts --no-commit && \
    echo "@openzeppelin/contracts/=lib/openzeppelin-contracts/contracts/" > remappings.txt && \
    forge config --evm-version cancun  && \
    forge test && \
    rm -r script src test
# 安装nodejs和npm
RUN apk add --no-cache nodejs npm

# 设置hardhat项目
WORKDIR /hardhat
RUN npm config set registry https://registry.npm.taobao.org && \
    echo -e "require(\"@nomicfoundation/hardhat-toolbox\");\n\nmodule.exports = {\n  solidity: \"0.8.17\",\n};" > hardhat.config.js && \
    npm init -y && \
    npm install --save-dev hardhat@2.14.0 @nomicfoundation/hardhat-toolbox@2.0.2 ethereum-waffle@4.0.10 @openzeppelin/contracts@4.7.1 && \
    npm install

ENTRYPOINT ["anvil"]