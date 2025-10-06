#!/bin/bash

export NAME="PixelDiff"
export PNAME="PixelDiff"
export VERSION="1.0.0"

rm -rf bin
mkdir -p bin
cp *.md bin/

export CGO_ENABLED=0

go generate

export GOOS=windows

echo "Compiling Windows x86"
export GOARCH=386
mkdir -p bin/${PNAME}_windows-x86
go build -o bin/${PNAME}_windows-x86/${NAME}.exe
cp *.md bin/${PNAME}_windows-x86/

echo "Compiling Windows x64"
export GOARCH=amd64
mkdir -p bin/${PNAME}_windows-x64
go build -o bin/${PNAME}_windows-x64/${NAME}.exe
cp *.md bin/${PNAME}_windows-x64/

echo "Compiling Windows ARM64"
export GOARCH=arm64
mkdir -p bin/${PNAME}_windows-arm64
go build -o bin/${PNAME}_windows-arm64/${NAME}.exe
cp *.md bin/${PNAME}_windows-arm64/

rm -f *.syso

export GOOS=darwin

echo "Compiling macOS x64"
export GOARCH=amd64
mkdir -p bin/${PNAME}_macos-x64
go build -o bin/${PNAME}_macos-x64/${NAME}
cp *.md bin/${PNAME}_macos-x64/
cp ico/icon.icns bin/${PNAME}_macos-x64/${NAME}.icns

echo "Compiling macOS ARM64"
export GOARCH=arm64
mkdir -p bin/${PNAME}_macos-arm64
go build -o bin/${PNAME}_macos-arm64/${NAME}
cp *.md bin/${PNAME}_macos-arm64/
cp ico/icon.icns bin/${PNAME}_macos-arm64/${NAME}.icns

export GOOS=linux

echo "Compiling Linux x86"
export GOARCH=386
mkdir -p bin/${PNAME}_linux-x86
go build -o bin/${PNAME}_linux-x86/${NAME}
cp *.md bin/${PNAME}_linux-x86/
cp ico/icon.png bin/${PNAME}_linux-x86/${NAME}.png

echo "Compiling Linux x64"
export GOARCH=amd64
mkdir -p bin/${PNAME}_linux-x64
go build -o bin/${PNAME}_linux-x64/${NAME}
cp *.md bin/${PNAME}_linux-x64/
cp ico/icon.png bin/${PNAME}_linux-x64/${NAME}.png

echo "Compiling Linux ARM32"
export GOARCH=arm
mkdir -p bin/${PNAME}_linux-arm32
go build -o bin/${PNAME}_linux-arm32/${NAME}
cp *.md bin/${PNAME}_linux-arm32/
cp ico/icon.png bin/${PNAME}_linux-arm32/${NAME}.png

echo "Compiling Linux ARM64"
export GOARCH=arm64
mkdir -p bin/${PNAME}_linux-arm64
go build -o bin/${PNAME}_linux-arm64/${NAME}
cp *.md bin/${PNAME}_linux-arm64/
cp ico/icon.png bin/${PNAME}_linux-arm64/${NAME}.png

rm -f bin/*.md

# 重置环境变量
unset NAME VERSION CGO_ENABLED GOOS GOARCH PNAME

echo "Compiling Local (Current OS)"
go build -o "$GOPATH/bin/$NAME"
rm -f *.syso
go clean
echo "Local build path: $GOPATH/bin/$NAME"
