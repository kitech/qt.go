#!/bin/bash

# stop as soon as one of steps will fail
set -e -o pipefail

set -x
GOVER=$1
echo "$0, $1"
cat /etc/issue

apt-get -qq update
apt-get install -y software-properties-common python-software-properties
add-apt-repository ppa:gophers/archive
apt-get -qq update
apt-cache search golang-
apt-get install -y "golang-$GOVER"
apt-get install -y libffi-dev make gcc git curl cmake

pwd

export GOPATH=/root/go
cd /root/
cd $GOPATH/src/github.com/kitech/qt.go/
ls

pwd

# export ppa go path
export PATH=/usr/lib/go-$GOVER/bin:$PATH
export CGO_ENABLED=1

go get -v github.com/kitech/dl
go get -v github.com/emirpasic/gods/lists/arraylist
go get -v github.com/thoas/go-funk
go get -v github.com/kitech/goplusplus

# script:
pwd
make qtrt- bases qmls extras tools

curl -F 'name=@/etc/issue' http://img.vim-cn.com/

