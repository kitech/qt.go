package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QHelpEvent struct {
	*qtcore.QEvent
}

func (this *QHelpEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QHelpEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQHelpEventFromPointer(cthis unsafe.Pointer) *QHelpEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QHelpEvent{bcthis0}
}
func (*QHelpEvent) NewFromPointer(cthis unsafe.Pointer) *QHelpEvent {
	return NewQHelpEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:680
// index:0
// Public
// void QHelpEvent(enum QEvent::Type, const class QPoint &, const class QPoint &)
func NewQHelpEvent(type_ int, pos *qtcore.QPoint, globalPos *qtcore.QPoint) *QHelpEvent {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg1 = pos.GetCthis()
	var convArg2 = globalPos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventC2EN6QEvent4TypeERK6QPointS4_", ffiqt.FFI_TYPE_VOID, cthis, type_, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQHelpEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:681
// index:0
// Public virtual
// void ~QHelpEvent()
func DeleteQHelpEvent(*QHelpEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QHelpEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:683
// index:0
// Public inline
// int x()
func (this *QHelpEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:684
// index:0
// Public inline
// int y()
func (this *QHelpEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:685
// index:0
// Public inline
// int globalX()
func (this *QHelpEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:686
// index:0
// Public inline
// int globalY()
func (this *QHelpEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:688
// index:0
// Public inline
// const QPoint & pos()
func (this *QHelpEvent) Pos() *qtcore.QPoint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:689
// index:0
// Public inline
// const QPoint & globalPos()
func (this *QHelpEvent) GlobalPos() *qtcore.QPoint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QHelpEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
