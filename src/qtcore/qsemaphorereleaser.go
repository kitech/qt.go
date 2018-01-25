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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSemaphoreReleaserFromPointer(cthis unsafe.Pointer) *QSemaphoreReleaser {
	return &QSemaphoreReleaser{&qtrt.CObject{cthis}}
}
func (*QSemaphoreReleaser) NewFromPointer(cthis unsafe.Pointer) *QSemaphoreReleaser {
	return NewQSemaphoreReleaserFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsemaphore.h:75
// index:0
// Public inline
// void QSemaphoreReleaser()
func NewQSemaphoreReleaser() *QSemaphoreReleaser {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:76
// index:1
// Public inline
// void QSemaphoreReleaser(class QSemaphore &, int)
func NewQSemaphoreReleaser_1(sem *QSemaphore, n int) *QSemaphoreReleaser {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2ER10QSemaphorei", ffiqt.FFI_TYPE_VOID, cthis, convArg0, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:78
// index:2
// Public inline
// void QSemaphoreReleaser(class QSemaphore *, int)
func NewQSemaphoreReleaser_2(sem *QSemaphore /*444 QSemaphore **/, n int) *QSemaphoreReleaser {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaserC2EP10QSemaphorei", ffiqt.FFI_TYPE_VOID, cthis, convArg0, n)
	gopp.ErrPrint(err, rv)
	gothis := NewQSemaphoreReleaserFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsemaphore.h:86
// index:0
// Public inline
// void ~QSemaphoreReleaser()
func DeleteQSemaphoreReleaser(*QSemaphoreReleaser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaserD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:92
// index:0
// Public inline
// void swap(class QSemaphoreReleaser &)
func (this *QSemaphoreReleaser) Swap(other *QSemaphoreReleaser) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaser4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsemaphore.h:98
// index:0
// Public inline
// QSemaphore * semaphore()
func (this *QSemaphoreReleaser) Semaphore() *QSemaphore /*444 QSemaphore **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSemaphoreReleaser9semaphoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsemaphore.h:101
// index:0
// Public inline
// QSemaphore * cancel()
func (this *QSemaphoreReleaser) Cancel() *QSemaphore /*444 QSemaphore **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSemaphoreReleaser6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSemaphoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
