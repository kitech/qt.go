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
// extern C begin: 4
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

/*

 */
type QCloseEvent struct {
	*qtcore.QEvent
}
type QCloseEvent_ITF interface {
	qtcore.QEvent_ITF
	QCloseEvent_PTR() *QCloseEvent
}

func (ptr *QCloseEvent) QCloseEvent_PTR() *QCloseEvent { return ptr }

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

// /usr/include/qt/QtGui/qevent.h:470
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCloseEvent()

/*

 */
func (*QCloseEvent) NewForInherit() *QCloseEvent {
	return NewQCloseEvent()
}
func NewQCloseEvent() *QCloseEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCloseEventC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCloseEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCloseEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:471
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCloseEvent()

/*

 */
func DeleteQCloseEvent(this *QCloseEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QCloseEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
