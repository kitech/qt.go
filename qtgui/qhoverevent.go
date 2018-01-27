package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QHoverEvent struct {
	*QInputEvent
}

func (this *QHoverEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QHoverEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQHoverEventFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QHoverEvent{bcthis0}
}
func (*QHoverEvent) NewFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	return NewQHoverEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:158
// index:0
// Public virtual
// void ~QHoverEvent()
func DeleteQHoverEvent(*QHoverEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QHoverEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:161
// index:0
// Public inline
// QPoint pos()
func (this *QHoverEvent) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:162
// index:0
// Public inline
// QPoint oldPos()
func (this *QHoverEvent) OldPos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent6oldPosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:165
// index:0
// Public inline
// const QPointF & posF()
func (this *QHoverEvent) PosF() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:166
// index:0
// Public inline
// const QPointF & oldPosF()
func (this *QHoverEvent) OldPosF() *qtcore.QPointF {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent7oldPosFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

//  body block end
