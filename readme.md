
### qt.go

Qt5 binding for Go (Golang) without CGO that aims to achieve Go's native compile speeds. Instead of using common bindings and heavy C++ wrapper code that forces you to compile and link time and time again, Qt.Go uses FFI so there's only a runtime dependency.

[![Build Status](https://travis-ci.org/kitech/qt.go.svg?branch=master)](https://travis-ci.org/kitech/qt.go)
[![Go Report Card](https://goreportcard.com/badge/github.com/kitech/qt.go)](https://goreportcard.com/report/github.com/kitech/qt.go)
[![GoDoc](https://godoc.org/github.com/kitech/qt.go?status.svg)](https://godoc.org/github.com/kitech/qt.go)
[![Sourcegraph](https://sourcegraph.com/github.com/kitech/qt.go/-/badge.svg)](https://sourcegraph.com/github.com/kitech/qt.go?badge)

### Features

* Binding code with no CGO compile cost
* Popular Qt5 packages (widgets/QML/extras) support
* Simple go-uic, go-rcc tools
* full signal/slot support
* protected method override support
* default arguments and value wrapper functions
* Class/Method/Function/Enum comment for godoc
* Go side signal/slot definition (experimental)


### Multiple platforms support
All platforms should be supported, for now some of them are tested:

* Archlinux/Ubuntu16+
* MacOS
* Android
* Windows

### Installation

##### requirement

* go 1.9+
* libffi
* dlfcn (windows)


##### FFI

Make sure libffi is installed

Debian based: `apt-get install libffi-dev`

Arch based: `pacman -S libffi`

MacOS: `brew install libffi`

##### qt.go:

    go get -v -u github.com/kitech/qt.go
    
##### runtime dependency:

    git clone https://github.com/kitech/qt.inline.git
    cd qt.inline
    cmake .
    make
    cp libQt5Inline.so /usr/lib/libQt5Inline.so

##### uic/rcc

    go get -v -u github.com/kitech/qt.go/cmd/go-uic
    go get -v -u github.com/kitech/qt.go/cmd/go-rcc

[Full Installation](https://github.com/kitech/qt.go/blob/master/install.md)

### Examples

    package main
    import "os"
    import "github.com/kitech/qt.go/qtwidgets"
    func main() {
        app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
        btn := qtwidgets.NewQPushButton1("hello qt.go", nil)
        btn.Show()
        app.Exec()
    }

More complex examples: https://github.com/kitech/qt.go/examples/ https://github.com/qtchina/qt.go.demos/ 

Go side signal/slot: [syntax document](https://github.com/kitech/qt.go/blob/master/docs/qt_meta_data_mark_syntax_for_go.md) [usage demo](https://github.com/kitech/qt.go/blob/master/qtmeta/tests/meta_data_test_.go)


### Community

  * QQ groupchat 933636020
  * Telegram room https://t.me/qtdevjiaoliu (Thanks https://github.com/xiayesuifeng)


### Internals

Qt.Go uses FFI to call wrapped Qt functions and methods, so there is no compile/link time dependency on Qt, only a run time dependency.

This should make the development and testing phases much faster.

[Internal document](https://github.com/kitech/qt.go/blob/master/docs/qt-go-internals.md)

