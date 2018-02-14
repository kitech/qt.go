package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QWhatsThisClickedEvent struct {
	*qtcore.QEvent
}
type QWhatsThisClickedEvent_ITF interface {
	qtcore.QEvent_ITF
	QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent
}

func (ptr *QWhatsThisClickedEvent) QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent { return ptr }

func (this *QWhatsThisClickedEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QWhatsThisClickedEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQWhatsThisClickedEventFromPointer(cthis unsafe.Pointer) *QWhatsThisClickedEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QWhatsThisClickedEvent{bcthis0}
}
func (*QWhatsThisClickedEvent) NewFromPointer(cthis unsafe.Pointer) *QWhatsThisClickedEvent {
	return NewQWhatsThisClickedEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:713
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWhatsThisClickedEvent(const QString &)
func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	var tmpArg0 = qtcore.NewQString_5(href)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWhatsThisClickedEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWhatsThisClickedEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:714
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWhatsThisClickedEvent()
func DeleteQWhatsThisClickedEvent(this *QWhatsThisClickedEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QWhatsThisClickedEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:716
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString href()
func (this *QWhatsThisClickedEvent) Href() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QWhatsThisClickedEvent4hrefEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
