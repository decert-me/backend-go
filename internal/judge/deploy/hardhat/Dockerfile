FROM node:18.16-alpine

# 环境配置
RUN apk add git && git config --global user.email "you@example.com" && git config --global user.name "Your Name" && \
    mkdir /hardhat && cd /hardhat &&  npm config set registry https://registry.npm.taobao.org && \
    echo -e "require(\"@nomicfoundation/hardhat-toolbox\");\n\nmodule.exports = {\n  solidity: \"0.8.17\",\n};" > hardhat.config.js && \
    npm init -y && npm install --save-dev hardhat@2.14.0 @nomicfoundation/hardhat-toolbox@2.0.2 && \
    npm install
