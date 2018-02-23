#!/bin/bash

# Fail out on error
set -e -o pipefail

# brew upgrade cmake
brew install make
brew install libffi

export PATH="/usr/local/opt/qt/bin:$PATH"

go get -v github.com/gonuts/dl
go get -v github.com/emirpasic/gods/lists/arraylist
go get -v github.com/thoas/go-funk
go get -v github.com/kitech/goplusplus

pwd
cd $GOPATH/src/github.com/kitech/qt.go

pwd
export CGO_ENABLED=1
export CGO_CFLAGS="-I/usr/local/opt/libffi/lib/libffi-3.2.1/include"
export CGO_LDFLAGS="-L/usr/local/opt/libffi/lib"
make qtrt- bases qmls extras tools

