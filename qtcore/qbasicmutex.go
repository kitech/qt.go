package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QBasicMutex struct {
	*qtrt.CObject
}
type QBasicMutex_ITF interface {
	QBasicMutex_PTR() *QBasicMutex
}

func (ptr *QBasicMutex) QBasicMutex_PTR() *QBasicMutex { return ptr }

func (this *QBasicMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBasicMutex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBasicMutexFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return &QBasicMutex{&qtrt.CObject{cthis}}
}
func (*QBasicMutex) NewFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return NewQBasicMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QBasicMutex()
func NewQBasicMutex() *QBasicMutex {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutexC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBasicMutexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBasicMutex)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void lock()
func (this *QBasicMutex) Lock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QBasicMutex) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool tryLock()
func (this *QBasicMutex) TryLock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex7tryLockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool try_lock()
func (this *QBasicMutex) Try_lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex8try_lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRecursive()
func (this *QBasicMutex) IsRecursive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:97
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isRecursive() const
func (this *QBasicMutex) IsRecursive_1() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QBasicMutex11isRecursiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQBasicMutex(this *QBasicMutex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicMutexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
