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
// extern C begin: 5
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
type QPaintEvent struct {
	*qtcore.QEvent
}
type QPaintEvent_ITF interface {
	qtcore.QEvent_ITF
	QPaintEvent_PTR() *QPaintEvent
}

func (ptr *QPaintEvent) QPaintEvent_PTR() *QPaintEvent { return ptr }

func (this *QPaintEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QPaintEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQPaintEventFromPointer(cthis unsafe.Pointer) *QPaintEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QPaintEvent{bcthis0}
}
func (*QPaintEvent) NewFromPointer(cthis unsafe.Pointer) *QPaintEvent {
	return NewQPaintEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:405
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPaintEvent(const QRegion &)

/*

 */
func NewQPaintEvent(paintRegion QRegion_ITF) *QPaintEvent {
	var convArg0 unsafe.Pointer
	if paintRegion != nil && paintRegion.QRegion_PTR() != nil {
		convArg0 = paintRegion.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPaintEventC2ERK7QRegion", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:406
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPaintEvent(const QRect &)

/*

 */
func NewQPaintEvent_1(paintRect qtcore.QRect_ITF) *QPaintEvent {
	var convArg0 unsafe.Pointer
	if paintRect != nil && paintRect.QRect_PTR() != nil {
		convArg0 = paintRect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPaintEventC2ERK5QRect", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPaintEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPaintEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:407
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPaintEvent()

/*

 */
func DeleteQPaintEvent(this *QPaintEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPaintEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:409
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QRect & rect() const

/*

 */
func (this *QPaintEvent) Rect() *qtcore.QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPaintEvent4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:410
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QRegion & region() const

/*

 */
func (this *QPaintEvent) Region() *QRegion {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPaintEvent6regionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
