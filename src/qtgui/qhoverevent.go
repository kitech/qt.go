//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QHoverEvent struct {
	*QInputEvent
}

func (this *QHoverEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}
func NewQHoverEventFromPointer(cthis unsafe.Pointer) *QHoverEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QHoverEvent{bcthis0}
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
func (this *QHoverEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:162
// index:0
// Public inline
// QPoint oldPos()
func (this *QHoverEvent) OldPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent6oldPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:165
// index:0
// Public inline
// const QPointF & posF()
func (this *QHoverEvent) PosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:166
// index:0
// Public inline
// const QPointF & oldPosF()
func (this *QHoverEvent) OldPosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QHoverEvent7oldPosFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
