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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:74
// index:0
// void QSystemSemaphore(const class QString &, int, enum QSystemSemaphore::AccessMode)
func NewQSystemSemaphore(key unsafe.Pointer, initialValue int, mode int) *QSystemSemaphore {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreC2ERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_VOID, cthis, key, &initialValue, &mode)
	gopp.ErrPrint(err, rv)
	return &QSystemSemaphore{cthis}
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:75
// index:0
// void ~QSystemSemaphore()
func DeleteQSystemSemaphore(*QSystemSemaphore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:77
// index:0
// void setKey(const class QString &, int, enum QSystemSemaphore::AccessMode)
func (this *QSystemSemaphore) SetKey(key unsafe.Pointer, initialValue int, mode int) {
	// 0: (, const QString & key, int initialValue, QSystemSemaphore::AccessMode mode), (key, &initialValue, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore6setKeyERK7QStringiNS_10AccessModeE", ffiqt.FFI_TYPE_VOID, this.cthis, key, &initialValue, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:78
// index:0
// QString key()
func (this *QSystemSemaphore) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:80
// index:0
// bool acquire()
func (this *QSystemSemaphore) Acquire() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7acquireEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:81
// index:0
// bool release(int)
func (this *QSystemSemaphore) Release(n int) {
	// 0: (, int n), (&n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSystemSemaphore7releaseEi", ffiqt.FFI_TYPE_VOID, this.cthis, &n)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:83
// index:0
// QSystemSemaphore::SystemSemaphoreError error()
func (this *QSystemSemaphore) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsystemsemaphore.h:84
// index:0
// QString errorString()
func (this *QSystemSemaphore) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSystemSemaphore11errorStringEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
