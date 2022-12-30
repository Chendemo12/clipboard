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
echo "SystemOS: linux"
echo "Platform: amd64"

echo "Building..."
echo ""

export CGO_ENABLED=0 # 禁用CGO
export GOOS=linux    # 目标平台是linux
export GOARCH=amd64  # 目标处理器架构是amd64

mkdir -p ./bin
# -s：忽略符号表和调试信息。
# -w：忽略DWARFv3调试信息，使用该选项后将无法使用gdb进行调试。
go build -tags=jsoniter -ldflags="-s -w" -gcflags='-l -l -l -m' ./main.go && mv main ./bin/"$FILENAME"

echo ""
echo "Finshed, output: ./bin/$FILENAME"
