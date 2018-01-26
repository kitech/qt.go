package qtqml

// /usr/include/qt/QtQml/qqmlparserstatus.h
// #include <qqmlparserstatus.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QQmlParserStatus struct {
	*qtrt.CObject
}

func (this *QQmlParserStatus) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlParserStatus) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlParserStatusFromPointer(cthis unsafe.Pointer) *QQmlParserStatus {
	return &QQmlParserStatus{&qtrt.CObject{cthis}}
}
func (*QQmlParserStatus) NewFromPointer(cthis unsafe.Pointer) *QQmlParserStatus {
	return NewQQmlParserStatusFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:52
// index:0
// Public
// void QQmlParserStatus()
func NewQQmlParserStatus() *QQmlParserStatus {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlParserStatusC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlParserStatusFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:53
// index:0
// Public virtual
// void ~QQmlParserStatus()
func DeleteQQmlParserStatus(*QQmlParserStatus) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlParserStatusD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:55
// index:0
// Public pure virtual
// void classBegin()
func (this *QQmlParserStatus) ClassBegin() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlParserStatus10classBeginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlparserstatus.h:56
// index:0
// Public pure virtual
// void componentComplete()
func (this *QQmlParserStatus) ComponentComplete() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QQmlParserStatus17componentCompleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
