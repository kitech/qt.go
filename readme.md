
### qt.go

Qt5 binding for Go (Golang) but without CGO aims get go's compile speed again. Instead of common binding using CGO and include heavy C++ wrapper code and compile and link these time and time again.

[![GoDoc](https://godoc.org/github.com/kitech/qt.go?status.svg)](https://godoc.org/github.com/kitech/qt.go)


### Features

* binding code no compile cost
* populate qt5 packages (widgets/qml/extras) support
* go-uic, go-rcc tools
* full signal/slot support
* protected method override support

### Installation

##### qt.go:

    go get -v -u github.com/kitech/qt.go
    

##### runtime dependency:

    git clone https://github.com/kitech/qt.inline.git
    cd qt.inline
    cmake .
    make
    cp libQtInline.so /usr/lib/libQt5Inline.so


### Examples

    package main
    import "os"
    import "github.com/qt.go/qtwidgets"
    func main() {
        app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
        btn := qtwidgets.NewQPushButton_1("hello qt.go", nil)
        btn.Show()
        app.Exec()
    }


### Internals

The main idea is using FFI to call wrapped qt functions/methods.

So there is no compile and link time dependency of qt lib, only runtime dependency.

When develop use qt.go, compile/test step can be very quickly.

