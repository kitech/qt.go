package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/qtrt"

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QReadLocker struct {
	*qtrt.CObject
}

func (this *QReadLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQReadLockerFromPointer(cthis unsafe.Pointer) *QReadLocker {
	return &QReadLocker{&qtrt.CObject{cthis}}
}
func (*QReadLocker) NewFromPointer(cthis unsafe.Pointer) *QReadLocker {
	return NewQReadLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qreadwritelock.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QReadLocker(QReadWriteLock *)
func NewQReadLocker(readWriteLock *QReadWriteLock /*777 QReadWriteLock **/) *QReadLocker {
	var convArg0 = readWriteLock.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLockerC2EP14QReadWriteLock", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQReadLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQReadLocker)
	return gothis
}

// /usr/include/qt/QtCore/qreadwritelock.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QReadLocker()
func DeleteQReadLocker(this *QReadLocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qreadwritelock.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QReadLocker) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLocker6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void relock()
func (this *QReadLocker) Relock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QReadLocker6relockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qreadwritelock.h:112
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QReadWriteLock * readWriteLock()
func (this *QReadLocker) ReadWriteLock() *QReadWriteLock /*777 QReadWriteLock **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QReadLocker13readWriteLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQReadWriteLockFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end
