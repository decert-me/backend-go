#!/bin/bash

osType=$(uname -s)

if [ "$osType" = "Linux" ]; then
    fileName="solc-static-linux"
elif [ "$osType" = "Darwin" ]; then
    fileName="solc-macos"
fi

VERSIONS=("v0.8.19")

for VERSION in "${VERSIONS[@]}"
do
    wget "https://github.com/ethereum/solidity/releases/download/$VERSION/$fileName"
    mv $fileName "solc-${VERSION:1}"
    mkdir -p "$HOME/.svm/${VERSION:1}"
    mv "solc-${VERSION:1}" "$HOME/.svm/${VERSION:1}/solc-${VERSION:1}"
    chmod +x "$HOME/.svm/${VERSION:1}/solc-${VERSION:1}"
done

#nohup anvil > /dev/null 2>&1 &

#./decert-judge > log.log 2>&1