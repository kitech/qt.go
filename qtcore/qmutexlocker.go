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
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QMutexLocker struct {
	*qtrt.CObject
}

func (this *QMutexLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMutexLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMutexLockerFromPointer(cthis unsafe.Pointer) *QMutexLocker {
	return &QMutexLocker{&qtrt.CObject{cthis}}
}
func (*QMutexLocker) NewFromPointer(cthis unsafe.Pointer) *QMutexLocker {
	return NewQMutexLockerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:199
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMutexLocker(QBasicMutex *)
func NewQMutexLocker(m *QBasicMutex /*777 QBasicMutex **/) *QMutexLocker {
	var convArg0 = m.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMutexLockerC2EP11QBasicMutex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMutexLockerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMutexLocker)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:213
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QMutexLocker()
func DeleteQMutexLocker(this *QMutexLocker) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMutexLockerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qmutex.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QMutexLocker) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMutexLocker6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void relock()
func (this *QMutexLocker) Relock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QMutexLocker6relockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QMutex * mutex()
func (this *QMutexLocker) Mutex() *QMutex /*777 QMutex **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QMutexLocker5mutexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMutexFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
		qtrt.KeepMe()
	}
}

//  keep block end
