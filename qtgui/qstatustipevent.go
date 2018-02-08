package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QStatusTipEvent struct {
	*qtcore.QEvent
}

func (this *QStatusTipEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QStatusTipEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQStatusTipEventFromPointer(cthis unsafe.Pointer) *QStatusTipEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QStatusTipEvent{bcthis0}
}
func (*QStatusTipEvent) NewFromPointer(cthis unsafe.Pointer) *QStatusTipEvent {
	return NewQStatusTipEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:700
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStatusTipEvent(const QString &)
func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	var tmpArg0 = qtcore.NewQString_5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStatusTipEventC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStatusTipEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStatusTipEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:701
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStatusTipEvent()
func DeleteQStatusTipEvent(this *QStatusTipEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QStatusTipEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:703
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString tip()
func (this *QStatusTipEvent) Tip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QStatusTipEvent3tipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

//  body block end
