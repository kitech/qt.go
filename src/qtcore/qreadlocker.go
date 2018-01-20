//  header block begin
// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QReadLocker struct {
	*qtrt.CObject
}

func (this *QReadLocker) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQReadLockerFromPointer(cthis unsafe.Pointer) *QReadLocker {
	return &QReadLocker{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qreadwritelock.h:87
// index:0
// Public inline
// void QReadLocker(class QReadWriteLock *)
func NewQReadLocker(readWriteLock unsafe.Pointer) *QReadLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QReadLockerC2EP14QReadWriteLock", ffiqt.FFI_TYPE_VOID, cthis, readWriteLock)
	gopp.ErrPrint(err, rv)
	gothis := NewQReadLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:89
// index:0
// Public inline
// void ~QReadLocker()
func DeleteQReadLocker(*QReadLocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QReadLockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:92
// index:0
// Public inline
// void unlock()
func (this *QReadLocker) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QReadLocker6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:102
// index:0
// Public inline
// void relock()
func (this *QReadLocker) Relock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QReadLocker6relockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:112
// index:0
// Public inline
// QReadWriteLock * readWriteLock()
func (this *QReadLocker) ReadWriteLock() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QReadLocker13readWriteLockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
