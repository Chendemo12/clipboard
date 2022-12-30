#!/usr/bin/env bash

# shellcheck disable=SC2002
PACKAGE_CONTROL=$(cat project/DEBIAN/control | sed ':jix;N;s/\n/ /g;b jix')

PACKAGE_NAME=${PACKAGE_CONTROL#*Package: }
PACKAGE_NAME=${PACKAGE_NAME%% *}

# 将第一个设置为打包后的文件名
FILENAME=$1
if [[ -z "$FILENAME" ]]; then
  FILENAME=${PACKAGE_NAME}
fi

echo ""
echo "--------------------------------"
echo ""
echo "Building project: $FILENAME"

echo ""
echo "SystemOS: Windows"
echo "Platform: amd64"

echo "Building..."
echo ""

export CGO_ENABLED=0
export GOOS=windows
export GOARCH=amd64

mkdir -p ./bin
go build -tags=jsoniter -ldflags="-s -w" -gcflags='-l -l -l -m' ./main.go && mv main.exe ./bin/"$FILENAME".exe

echo ""
echo "Finshed, output: ./bin/$FILENAME.exe"
