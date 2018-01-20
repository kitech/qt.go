//  header block begin
// /usr/include/qt/QtCore/qsystemsemaphore.h
// #include <qsystemsemaphore.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QSystemSemaphore struct {
	*qtrt.CObject
}

func (this *QSystemSemaphore) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQSystemSemaphoreFromPointer(cthis unsafe.Pointer) *QSystemSemaphore {
	return &QSystemSemaphore{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// Public
// void QSystemSemaphore(const class QString &, int, enum QSystemSemaphore::AccessMode)
func NewQSystemSemaphore(key *QString, initialValue int, mode int) *QSystemSemaphore {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &initialValue, &mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQSystemSemaphoreFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:75
// index:0
// Public
// void ~QSystemSemaphore()
func DeleteQSystemSemaphore(*QSystemSemaphore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// Public
// void setKey(const class QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey(key *QString, initialValue int, mode int) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &initialValue, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:78
// index:0
// Public
// QString key()
func (this *QSystemSemaphore) Key() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:80
// index:0
// Public
// bool acquire()
func (this *QSystemSemaphore) Acquire() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7acquireEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// Public
// bool release(int)
func (this *QSystemSemaphore) Release(n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &n)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:83
// index:0
// Public
// QSystemSemaphore::SystemSemaphoreError error()
func (this *QSystemSemaphore) Error() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:84
// index:0
// Public
// QString errorString()
func (this *QSystemSemaphore) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
