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
type QEnterEvent struct {
	*qtcore.QEvent
}

func (this *QEnterEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QEnterEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQEnterEventFromPointer(cthis unsafe.Pointer) *QEnterEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QEnterEvent{bcthis0}
}
func (*QEnterEvent) NewFromPointer(cthis unsafe.Pointer) *QEnterEvent {
	return NewQEnterEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:85
// index:0
// Public
// void QEnterEvent(const QPointF &, const QPointF &, const QPointF &)
func NewQEnterEvent(localPos *qtcore.QPointF, windowPos *qtcore.QPointF, screenPos *qtcore.QPointF) *QEnterEvent {
	cthis := qtrt.Calloc(1, 256) // 72
	var convArg0 = localPos.GetCthis()
	var convArg1 = windowPos.GetCthis()
	var convArg2 = screenPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QEnterEventC2ERK7QPointFS2_S2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQEnterEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:86
// index:0
// Public virtual
// void ~QEnterEvent()
func DeleteQEnterEvent(*QEnterEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QEnterEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:89
// index:0
// Public inline
// QPoint pos()
func (this *QEnterEvent) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:90
// index:0
// Public inline
// QPoint globalPos()
func (this *QEnterEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:91
// index:0
// Public inline
// int x()
func (this *QEnterEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:92
// index:0
// Public inline
// int y()
func (this *QEnterEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:93
// index:0
// Public inline
// int globalX()
func (this *QEnterEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:94
// index:0
// Public inline
// int globalY()
func (this *QEnterEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:96
// index:0
// Public inline
// const QPointF & localPos()
func (this *QEnterEvent) LocalPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent8localPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:97
// index:0
// Public inline
// const QPointF & windowPos()
func (this *QEnterEvent) WindowPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9windowPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:98
// index:0
// Public inline
// const QPointF & screenPos()
func (this *QEnterEvent) ScreenPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QEnterEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
