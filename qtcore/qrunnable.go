package qtcore

// /usr/include/qt/QtCore/qrunnable.h
// #include <qrunnable.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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

type QRunnable struct {
	*qtrt.CObject
}

func (this *QRunnable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRunnable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRunnableFromPointer(cthis unsafe.Pointer) *QRunnable {
	return &QRunnable{&qtrt.CObject{cthis}}
}
func (*QRunnable) NewFromPointer(cthis unsafe.Pointer) *QRunnable {
	return NewQRunnableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrunnable.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void run()
func (this *QRunnable) Run() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnable3runEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrunnable.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRunnable()
func NewQRunnable() *QRunnable {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnableC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRunnableFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRunnable)
	return gothis
}

// /usr/include/qt/QtCore/qrunnable.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRunnable()
func DeleteQRunnable(this *QRunnable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qrunnable.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool autoDelete()
func (this *QRunnable) AutoDelete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QRunnable10autoDeleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrunnable.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAutoDelete(_Bool)
func (this *QRunnable) SetAutoDelete(_autoDelete bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnable13setAutoDeleteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _autoDelete)
	gopp.ErrPrint(err, rv)
}

//  body block end
