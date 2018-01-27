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
// Public
// void QStatusTipEvent(const QString &)
func NewQStatusTipEvent(tip *qtcore.QString) *QStatusTipEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStatusTipEventC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStatusTipEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:701
// index:0
// Public virtual
// void ~QStatusTipEvent()
func DeleteQStatusTipEvent(*QStatusTipEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QStatusTipEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:703
// index:0
// Public inline
// QString tip()
func (this *QStatusTipEvent) Tip() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QStatusTipEvent3tipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
