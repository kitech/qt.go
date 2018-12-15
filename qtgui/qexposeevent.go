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
type QExposeEvent struct {
	*qtcore.QEvent
}
type QExposeEvent_ITF interface {
	qtcore.QEvent_ITF
	QExposeEvent_PTR() *QExposeEvent
}

func (ptr *QExposeEvent) QExposeEvent_PTR() *QExposeEvent { return ptr }

func (this *QExposeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QExposeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQExposeEventFromPointer(cthis unsafe.Pointer) *QExposeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QExposeEvent{bcthis0}
}
func (*QExposeEvent) NewFromPointer(cthis unsafe.Pointer) *QExposeEvent {
	return NewQExposeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:439
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QExposeEvent(const QRegion &)

/*

 */
func (*QExposeEvent) NewForInherit(rgn QRegion_ITF) *QExposeEvent {
	return NewQExposeEvent(rgn)
}
func NewQExposeEvent(rgn QRegion_ITF) *QExposeEvent {
	var convArg0 unsafe.Pointer
	if rgn != nil && rgn.QRegion_PTR() != nil {
		convArg0 = rgn.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QExposeEventC2ERK7QRegion", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQExposeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQExposeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:440
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QExposeEvent()

/*

 */
func DeleteQExposeEvent(this *QExposeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QExposeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:442
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QRegion & region() const

/*

 */
func (this *QExposeEvent) Region() *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QExposeEvent6regionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
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
