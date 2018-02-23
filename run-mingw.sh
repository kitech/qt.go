#!/bin/bash

# stop as soon as one of steps will fail
set -e -o pipefail

set -x
WINARCH=$1
echo "$0, $1"
cat /etc/issue

echo "[ownstuff]" >> /etc/pacman.conf
echo "SigLevel = Optional TrustAll" >> /etc/pacman.conf
echo "Server = https://martchus.no-ip.biz/repo/arch/ownstuff/os/\$arch" >> /etc/pacman.conf

pacman -Suyy
pacman -S --noconfirm go git make wget
pacman -S --noconfirm mingw-w64-gcc mingw-w64-libffi

pwd

export GOPATH=/root/go
cd /root/
cd $GOPATH/src/github.com/kitech/qt.go/
ls

pwd



git clone https://github.com/gonuts/dl.git
sed -i 's/NoLoad/\/\/NoLoad/' dl/dl.go
sed -i 's/NoDelete/\/\/NoDelete/' dl/dl.go
mkdir $GOPATH/src/github.com/gonuts -p
cp -a dl $GOPATH/src/github.com/gonuts/

export CGO_ENABLED=1
export GOOS=windows
if [ x"$WINARCH" = x"x64" ]; then
    ### build x64 version
    wget http://pkg.mxe.cc/repos/tar/mxe-x86-64-w64-mingw32.static/mxe-x86-64-w64-mingw32.static-dlfcn-win32_0.e19bf07.tar.xz
    tar xvf mxe-x86-64-w64-mingw32.static-dlfcn-win32_0.e19bf07.tar.xz
    cp -va usr/x86_64-w64-mingw32.static/* /usr/x86_64-w64-mingw32/

    export GOARCH=amd64
    export CC=x86_64-w64-mingw32-gcc
    export CGO_LDFLAGS="-L/usr/x86_64-w64-mingw32/lib/ -ldl -lpsapi -lkernel32"
    export CGO_CFLAGS="-I/usr/x86_64-w64-mingw32/lib/libffi-3.2.1/include/"

    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools

else
    ### build x32 version dll
    wget http://pkg.mxe.cc/repos/tar/mxe-i686-w64-mingw32.static/mxe-i686-w64-mingw32.static-dlfcn-win32_0.e19bf07.tar.xz
    tar xvf mxe-i686-w64-mingw32.static-dlfcn-win32_0.e19bf07.tar.xz
    cp -va usr/i686-w64-mingw32.static/* /usr/i686-w64-mingw32/

    export GOARCH=386
    export CC=i686-w64-mingw32-gcc
    export CGO_LDFLAGS="-L/usr/i686-w64-mingw32/lib/ -ldl -lpsapi -lkernel32"
    export CGO_CFLAGS="-I/usr/i686-w64-mingw32/lib/libffi-3.2.1/include/"

    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools
fi


curl -F 'name=@/etc/issue' http://img.vim-cn.com/

