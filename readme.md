
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

##### qt.go:

    go get -v -u github.com/kitech/qt.go
    
##### runtime dependency:

    git clone https://github.com/kitech/qt.inline.git
    cd qt.inline
    cmake .
    make
    cp libQt5Inline.so /usr/lib/libQt5Inline.so

##### FFI

Make sure libffi is installed
    
Debian based:

    apt-get install libffi-dev
    
Arch based:

    pacman -S libffi

[Full Installation](https://github.com/kitech/qt.go/blob/master/install.md)

### Examples

    package main
    import "os"
    import "github.com/kitech/qt.go/qtwidgets"
    func main() {
        app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
        btn := qtwidgets.NewQPushButton_1("hello qt.go", nil)
        btn.Show()
        app.Exec()
    }

### Internals

Qt.Go uses FFI to call wrapped Qt functions and methods, so there is no compile/link time dependency on Qt, only a run time dependency.

This should make the development and testing phases much faster.
