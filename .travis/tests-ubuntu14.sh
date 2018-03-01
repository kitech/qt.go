#!/bin/bash

# Fail out on error
set -e -o pipefail

./.travis/build-ubuntu14.sh

pwd
wget https://github.com/kitech/qt.inline/releases/download/v1.0-rc.1/qt510_linux_x64_static.tar.bz2

tar xvf qt510_linux_x64_static.tar.bz2

set -x
export LD_LIBRARY_PATH=$PWD/qt510_linux_x64_static
pwd
ls

sudo apt-get install -y libx11-6 libxcb-icccm4 libxcb-image0 libxcb-keysyms1 libxcb-present0 libxcb-present0 libxcb-shm0 libxcb-sync1 libxcb-xfixes0 libxcb-xkb1 libxcb-cursor0 libxcb-xinerama0

ldd $PWD/qt510_linux_x64_static/libQt5Inline.so
go test -v tests/hello_test.go

