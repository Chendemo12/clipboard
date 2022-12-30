#!/bin/bash

# shellcheck disable=SC2002
PACKAGE_CONTROL=$(cat project/DEBIAN/control | sed ':jix;N;s/\n/ /g;b jix')
PACKAGE_NAME=${PACKAGE_CONTROL#*Package: }
PACKAGE_NAME=${PACKAGE_NAME%% *}

PACKAGE_VERSION=$(cat src/application.go | sed ':jix;N;s/\n/ /g;b jix')
PACKAGE_VERSION=${PACKAGE_VERSION#*Version *\"}

PACKAGE_VERSION=${PACKAGE_VERSION%% *}
PACKAGE_VERSION=${PACKAGE_VERSION%%\"*}

# 修改版本号
sed -i "/^Version/s/ .*/ $PACKAGE_VERSION/" project/DEBIAN/control

echo "**********************************************"
echo "*"
echo "* Package Name: ${PACKAGE_NAME}"
echo "* Version: ${PACKAGE_VERSION}"
echo "*"
echo ""

if [ -e './tmp' ]; then
  rm -rf ./tmp
fi

echo "Building binary package..."
rm -rf ./bin
chmod -R 775 build/*
./build/linux.amd64.sh "${PACKAGE_NAME}" || exit 0
echo "Done."

mkdir ./tmp

rsync -av ./project ./tmp/ --exclude=.idea --exclude=logs
mkdir -p ./tmp/project/opt/cowave/bin
mv ./bin/"${PACKAGE_NAME}" ./tmp/project/opt/cowave/bin/"$PACKAGE_NAME"

# 确保拥有可执行权限
chmod -R 775 ./tmp/project/DEBIAN/*
chmod -R 775 ./tmp/project/opt/cowave/bin*

mkdir -p ./deb

# 强制使用gzip打包deb
# Ubuntu >= 21.04 Supported
dpkg-deb -Zxz --build --root-owner-group tmp/project deb/"${PACKAGE_NAME}"_"${PACKAGE_VERSION}"_amd64.deb

rm -rf ./tmp

exit 0
