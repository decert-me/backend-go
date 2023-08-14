FROM ghcr.io/foundry-rs/foundry:nightly-49080c1424987dd772db73504dbfa575de3cab14


# 环境配置
RUN apk update && apk add --no-cache git bash && git config --global user.email "you@example.com" && git config --global user.name "Your Name" && \
    mkdir /foundry && cd /foundry && forge init && forge test && \
    rm -r script src test

RUN apk add --no-cache nodejs && apk add --no-cache npm && \
    mkdir /hardhat && cd /hardhat &&  npm config set registry https://registry.npm.taobao.org && \
    echo -e "require(\"@nomicfoundation/hardhat-toolbox\");\n\nmodule.exports = {\n  solidity: \"0.8.17\",\n};" > hardhat.config.js && \
    npm init -y && npm install --save-dev hardhat@2.14.0 @nomicfoundation/hardhat-toolbox@2.0.2 ethereum-waffle@4.0.10 && \
    npm install

ENTRYPOINT ["anvil"]