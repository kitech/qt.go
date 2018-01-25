package qtcore

// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QEventLoopLocker struct {
	*qtrt.CObject
}

func (this *QEventLoopLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QEventLoopLocker) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQEventLoopLockerFromPointer(cthis unsafe.Pointer) *QEventLoopLocker {
	return &QEventLoopLocker{&qtrt.CObject{cthis}}
}
func (*QEventLoopLocker) NewFromPointer(cthis unsafe.Pointer) *QEventLoopLocker {
	return NewQEventLoopLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeventloop.h:93
// index:0
// Public
// void QEventLoopLocker()
func NewQEventLoopLocker() *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:94
// index:1
// Public
// void QEventLoopLocker(class QEventLoop *)
func NewQEventLoopLocker_1(loop *QEventLoop /*444 QEventLoop **/) *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = loop.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP10QEventLoop", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:95
// index:2
// Public
// void QEventLoopLocker(class QThread *)
func NewQEventLoopLocker_2(thread *QThread /*444 QThread **/) *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = thread.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP7QThread", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:96
// index:0
// Public
// void ~QEventLoopLocker()
func DeleteQEventLoopLocker(*QEventLoopLocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
