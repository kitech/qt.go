#!/bin/bash

# Fail out on error
set -e -o pipefail

./.travis/build-osx.sh

pwd
wget https://github.com/kitech/qt.inline/releases/download/v1.0-rc2/qt510_macos_x64_xcode8.3.tar.bz2

tar xvf qt510_macos_x64_xcode8.3.tar.bz2

set -x
brew install qt5

export LD_LIBRARY_PATH=$PWD/qt510_macos_x64_xcode8.3:/usr/local/opt/qt/lib
pwd
ls

cd $GOPATH/src/github.com/kitech/qt.go

pwd
export CGO_ENABLED=1
export CGO_CFLAGS="-I/usr/local/opt/libffi/lib/libffi-3.2.1/include"
export CGO_LDFLAGS="-L/usr/local/opt/libffi/lib"

file $PWD/qt510_macos_x64_xcode8.3/libQt5Inline.dylib
otool -L $PWD/qt510_macos_x64_xcode8.3/libQt5Inline.dylib

go test -v tests/hello_test.go

