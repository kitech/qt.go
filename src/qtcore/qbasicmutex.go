package qtcore

// /usr/include/qt/QtCore/qmutex.h
// #include <qmutex.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QBasicMutex struct {
	*qtrt.CObject
}

func (this *QBasicMutex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBasicMutex) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQBasicMutexFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return &QBasicMutex{&qtrt.CObject{cthis}}
}
func (*QBasicMutex) NewFromPointer(cthis unsafe.Pointer) *QBasicMutex {
	return NewQBasicMutexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmutex.h:71
// index:0
// Public inline
// void QBasicMutex()
func NewQBasicMutex() *QBasicMutex {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutexC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBasicMutexFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmutex.h:77
// index:0
// Public inline
// void lock()
func (this *QBasicMutex) Lock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutex4lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:83
// index:0
// Public inline
// void unlock()
func (this *QBasicMutex) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutex6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmutex.h:89
// index:0
// Public inline
// bool tryLock()
func (this *QBasicMutex) TryLock() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutex7tryLockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:94
// index:0
// Public inline
// bool try_lock()
func (this *QBasicMutex) Try_lock() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutex8try_lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:96
// index:0
// Public
// bool isRecursive()
func (this *QBasicMutex) IsRecursive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicMutex11isRecursiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmutex.h:97
// index:1
// Public
// bool isRecursive()
func (this *QBasicMutex) IsRecursive_1() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QBasicMutex11isRecursiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
