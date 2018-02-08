package qtcore

// /usr/include/qt/QtCore/qsemaphore.h
// #include <qsemaphore.h>
// #include <QtCore>

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

type QSemaphoreReleaser struct {
	*qtrt.CObject
}

func (this *QSemaphoreReleaser) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSemaphoreReleaser) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSemaphoreReleaserFromPointer(cthis unsafe.Pointer) *QSemaphoreReleaser {
	return &QSemaphoreReleaser{&qtrt.CObject{cthis}}
}
func (*QSemaphoreReleaser) NewFromPointer(cthis unsafe.Pointer) *QSemaphoreReleaser {
	return NewQSemaphoreReleaserFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsemaphore.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser()
func NewQSemaphoreReleaser() *QSemaphoreReleaser {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:76
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore &, int)
func NewQSemaphoreReleaser_1(sem *QSemaphore, n int) *QSemaphoreReleaser {
	var convArg0 = sem.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2ER10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:78
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore *, int)
func NewQSemaphoreReleaser_2(sem *QSemaphore /*777 QSemaphore **/, n int) *QSemaphoreReleaser {
	var convArg0 = sem.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2EP10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QSemaphoreReleaser()
func DeleteQSemaphoreReleaser(this *QSemaphoreReleaser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsemaphore.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSemaphoreReleaser &)
func (this *QSemaphoreReleaser) Swap(other *QSemaphoreReleaser) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaser4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSemaphore * semaphore()
func (this *QSemaphoreReleaser) Semaphore() *QSemaphore /*777 QSemaphore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSemaphoreReleaser9semaphoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsemaphore.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSemaphore * cancel()
func (this *QSemaphoreReleaser) Cancel() *QSemaphore /*777 QSemaphore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaser6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end
