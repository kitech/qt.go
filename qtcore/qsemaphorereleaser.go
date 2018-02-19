package qtcore

// /usr/include/qt/QtCore/qsemaphore.h
// #include <qsemaphore.h>
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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QSemaphoreReleaser struct {
	*qtrt.CObject
}
type QSemaphoreReleaser_ITF interface {
	QSemaphoreReleaser_PTR() *QSemaphoreReleaser
}

func (ptr *QSemaphoreReleaser) QSemaphoreReleaser_PTR() *QSemaphoreReleaser { return ptr }

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
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:76
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore &, int)
func NewQSemaphoreReleaser_1(sem QSemaphore_ITF, n int) *QSemaphoreReleaser {
	var convArg0 unsafe.Pointer
	if sem != nil && sem.QSemaphore_PTR() != nil {
		convArg0 = sem.QSemaphore_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2ER10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:76
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore &, int)
func NewQSemaphoreReleaser_1_(sem QSemaphore_ITF) *QSemaphoreReleaser {
	var convArg0 unsafe.Pointer
	if sem != nil && sem.QSemaphore_PTR() != nil {
		convArg0 = sem.QSemaphore_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2ER10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:78
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore *, int)
func NewQSemaphoreReleaser_2(sem QSemaphore_ITF /*777 QSemaphore **/, n int) *QSemaphoreReleaser {
	var convArg0 unsafe.Pointer
	if sem != nil && sem.QSemaphore_PTR() != nil {
		convArg0 = sem.QSemaphore_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2EP10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSemaphoreReleaser)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:78
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QSemaphoreReleaser(QSemaphore *, int)
func NewQSemaphoreReleaser_2_(sem QSemaphore_ITF /*777 QSemaphore **/) *QSemaphoreReleaser {
	var convArg0 unsafe.Pointer
	if sem != nil && sem.QSemaphore_PTR() != nil {
		convArg0 = sem.QSemaphore_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	n := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2EP10QSemaphorei", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsemaphore.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSemaphoreReleaser &)
func (this *QSemaphoreReleaser) Swap(other QSemaphoreReleaser_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSemaphoreReleaser_PTR() != nil {
		convArg0 = other.QSemaphoreReleaser_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaser4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:98
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSemaphore * semaphore() const
func (this *QSemaphoreReleaser) Semaphore() *QSemaphore /*777 QSemaphore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSemaphoreReleaser9semaphoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsemaphore.h:101
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSemaphore * cancel()
func (this *QSemaphoreReleaser) Cancel() *QSemaphore /*777 QSemaphore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSemaphoreReleaser6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
