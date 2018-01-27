package qtqml

// /usr/include/qt/QtQml/qqmlinfo.h
// #include <qqmlinfo.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlInfo struct {
	*qtrt.CObject
}

func (this *QQmlInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlInfo) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlInfoFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return &QQmlInfo{&qtrt.CObject{cthis}}
}
func (*QQmlInfo) NewFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return NewQQmlInfoFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlinfo.h:76
// index:0
// Public
// void ~QQmlInfo()
func DeleteQQmlInfo(*QQmlInfo) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QQmlInfoD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
