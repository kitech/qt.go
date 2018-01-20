//  header block begin
// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>
package qtcore

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
type QMutexLocker struct {
	*qtrt.CObject
}

func (this *QMutexLocker) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qmutex.h:199
// index:0
// inline
// void QMutexLocker(class QBasicMutex *)
func NewQMutexLocker(m unsafe.Pointer) *QMutexLocker {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLockerC2EP11QBasicMutex", ffiqt.FFI_TYPE_VOID, cthis, m)
	gopp.ErrPrint(err, rv)
	gothis := NewQMutexLockerFromPointer(cthis)
	return gothis
}
func NewQMutexLockerFromPointer(cthis unsafe.Pointer) *QMutexLocker {
	return &QMutexLocker{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qmutex.h:213
// index:0
// inline
// void ~QMutexLocker()
func DeleteQMutexLocker(*QMutexLocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:215
// index:0
// inline
// void unlock()
func (this *QMutexLocker) Unlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLocker6unlockEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:223
// index:0
// inline
// void relock()
func (this *QMutexLocker) Relock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMutexLocker6relockEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:238
// index:0
// inline
// QMutex * mutex()
func (this *QMutexLocker) Mutex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMutexLocker5mutexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
