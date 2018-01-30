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
type QExposeEvent struct {
	*qtcore.QEvent
}

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

// /usr/include/qt/QtGui/qevent.h:434
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QExposeEvent(const QRegion &)
func NewQExposeEvent(rgn *QRegion) *QExposeEvent {
	var convArg0 = rgn.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventC2ERK7QRegion", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQExposeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:435
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QExposeEvent()
func DeleteQExposeEvent(*QExposeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QExposeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:437
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QRegion & region()
func (this *QExposeEvent) Region() *QRegion {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QExposeEvent6regionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
