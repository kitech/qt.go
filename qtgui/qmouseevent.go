package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
type QMouseEvent struct {
	*QInputEvent
}

func (this *QMouseEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QMouseEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQMouseEventFromPointer(cthis unsafe.Pointer) *QMouseEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QMouseEvent{bcthis0}
}
func (*QMouseEvent) NewFromPointer(cthis unsafe.Pointer) *QMouseEvent {
	return NewQMouseEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:118
// index:0
// Public virtual
// void ~QMouseEvent()
func DeleteQMouseEvent(*QMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:121
// index:0
// Public inline
// QPoint pos()
func (this *QMouseEvent) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:122
// index:0
// Public inline
// QPoint globalPos()
func (this *QMouseEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:123
// index:0
// Public inline
// int x()
func (this *QMouseEvent) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:124
// index:0
// Public inline
// int y()
func (this *QMouseEvent) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:125
// index:0
// Public inline
// int globalX()
func (this *QMouseEvent) GlobalX() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:126
// index:0
// Public inline
// int globalY()
func (this *QMouseEvent) GlobalY() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qevent.h:128
// index:0
// Public inline
// const QPointF & localPos()
func (this *QMouseEvent) LocalPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent8localPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:129
// index:0
// Public inline
// const QPointF & windowPos()
func (this *QMouseEvent) WindowPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9windowPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:130
// index:0
// Public inline
// const QPointF & screenPos()
func (this *QMouseEvent) ScreenPos() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9screenPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:132
// index:0
// Public inline
// Qt::MouseButton button()
func (this *QMouseEvent) Button() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:133
// index:0
// Public inline
// Qt::MouseButtons buttons()
func (this *QMouseEvent) Buttons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:135
// index:0
// Public inline
// void setLocalPos(const class QPointF &)
func (this *QMouseEvent) SetLocalPos(localPosition *qtcore.QPointF) {
	var convArg0 = localPosition.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEvent11setLocalPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:141
// index:0
// Public
// Qt::MouseEventSource source()
func (this *QMouseEvent) Source() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6sourceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:142
// index:0
// Public
// Qt::MouseEventFlags flags()
func (this *QMouseEvent) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
