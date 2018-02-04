package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QCloseEvent struct {
	*qtcore.QEvent
}

func (this *QCloseEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QCloseEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQCloseEventFromPointer(cthis unsafe.Pointer) *QCloseEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QCloseEvent{bcthis0}
}
func (*QCloseEvent) NewFromPointer(cthis unsafe.Pointer) *QCloseEvent {
	return NewQCloseEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:477
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCloseEvent()
func NewQCloseEvent() *QCloseEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QCloseEventC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCloseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCloseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:478
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCloseEvent()
func DeleteQCloseEvent(this *QCloseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QCloseEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
