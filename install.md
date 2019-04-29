

### windows (static qt)

directories: all put in D:\qtenv\

##### gcc

install from mingw-64 i686-win32-sjlj-7.2.0

and make sure gcc is in %path%

    gcc -v

##### libffi && dlfcn && Qt5Inline.dll

    git clone https://github.com/qtchina/qt512_qtenv_win32.git D:/qtenv

##### qt.go

open cmd.exe:

    set CGO_CFLAGS=-ID:/qtenv/include -ID:/qtenv/lib/libffi-3.2.1/include/
    set CGO_LDFLAGS=-LD:/qtenv/lib
    
    go get github.com/kitech/dl
    go get github.com/kitech/qt.go

    cd D:/qtenv
    go build -o bin/button.exe %GOPATH%/src/github.com/kitech/qt.go/examples/button.go

    bin\button.exe


### macOS

directories: all put in $HOME/qtenv

xcode8.3 or xcode 9.3 are fine.

##### libffi && dlfcn && qt5

    brew install libffi  qt5

##### Qt5Inline.dylib

wget https://github.com/kitech/qt.inline/releases/download/v5.12.107/qt512_macos_x64_xcode9.3.tar.bz2

extract to $HOME/qtenv/libQt5Inline.dylib

##### qt.go

    export CGO_CFLAGS="-I/usr/local/opt/libffi/lib/libffi-3.2.1/include"
    export CGO_LDFLAGS="-L/usr/local/opt/libffi/lib"
    
    go get github.com/kitech/dl
    go get github.com/kitech/qt.go

    go build $GOPATH/src/github.com/kitech/qt.go/examples/button.go

    LD_LIBRARY_PATH=$HOME/qtenv:/usr/local/opt/qt/lib ./button 


### Linux (static qt)

##### libffi

* archlinux: `pacman -S libffi`
* ubuntu/debian: `apt-get install libffi-dev`

##### Qt5Inline.so

wget https://github.com/kitech/qt.inline/releases/download/v5.12.107/qt512_linux_x64_static.tar.bz2

extract to $HOME/qtenv/libQt5Inline.so

Note: thus not need install Qt yourself.

##### qt.go

    go get github.com/kitech/dl
    go get github.com/kitech/qt.go

    go build $GOPATH/src/github.com/kitech/qt.go/examples/button.go

    LD_LIBRARY_PATH=$HOME/qtenv ./button 


### notes

For windows/linux, the installation manual include static qt version Qt5Inline.

When previewed, and want deploy your applications, build yourself version Qt library.


