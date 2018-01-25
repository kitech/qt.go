package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QPaintEvent struct {
	*qtcore.QEvent
}

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
// Public
// void QPaintEvent(const class QRegion &)
func NewQPaintEvent(paintRegion *QRegion) *QPaintEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg0 = paintRegion.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventC2ERK7QRegion", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:406
// index:1
// Public
// void QPaintEvent(const class QRect &)
func NewQPaintEvent_1(paintRect *qtcore.QRect) *QPaintEvent {
	cthis := qtrt.Calloc(1, 256) // 56
	var convArg0 = paintRect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventC2ERK5QRect", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPaintEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:407
// index:0
// Public virtual
// void ~QPaintEvent()
func DeleteQPaintEvent(*QPaintEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPaintEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:409
// index:0
// Public inline
// const QRect & rect()
func (this *QPaintEvent) Rect() *qtcore.QRect {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPaintEvent4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:410
// index:0
// Public inline
// const QRegion & region()
func (this *QPaintEvent) Region() *QRegion {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPaintEvent6regionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
