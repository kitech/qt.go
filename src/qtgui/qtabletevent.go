//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QTabletEvent struct {
	*QInputEvent
}

func (this *QTabletEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}
func NewQTabletEventFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTabletEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:259
// index:0
// Public virtual
// void ~QTabletEvent()
func DeleteQTabletEvent(*QTabletEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTabletEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:261
// index:0
// Public inline
// QPoint pos()
func (this *QTabletEvent) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:262
// index:0
// Public inline
// QPoint globalPos()
func (this *QTabletEvent) GlobalPos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent9globalPosEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:267
// index:0
// Public inline
// const QPointF & posF()
func (this *QTabletEvent) PosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent4posFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:268
// index:0
// Public inline
// const QPointF & globalPosF()
func (this *QTabletEvent) GlobalPosF() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent10globalPosFEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:270
// index:0
// Public inline
// int x()
func (this *QTabletEvent) X() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:271
// index:0
// Public inline
// int y()
func (this *QTabletEvent) Y() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:272
// index:0
// Public inline
// int globalX()
func (this *QTabletEvent) GlobalX() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7globalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:273
// index:0
// Public inline
// int globalY()
func (this *QTabletEvent) GlobalY() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7globalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:274
// index:0
// Public inline
// qreal hiResGlobalX()
func (this *QTabletEvent) HiResGlobalX() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:275
// index:0
// Public inline
// qreal hiResGlobalY()
func (this *QTabletEvent) HiResGlobalY() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:276
// index:0
// Public inline
// QTabletEvent::TabletDevice device()
func (this *QTabletEvent) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:277
// index:0
// Public inline
// QTabletEvent::PointerType pointerType()
func (this *QTabletEvent) PointerType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent11pointerTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:278
// index:0
// Public inline
// qint64 uniqueId()
func (this *QTabletEvent) UniqueId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8uniqueIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:279
// index:0
// Public inline
// qreal pressure()
func (this *QTabletEvent) Pressure() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8pressureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:280
// index:0
// Public inline
// int z()
func (this *QTabletEvent) Z() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1zEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:281
// index:0
// Public inline
// qreal tangentialPressure()
func (this *QTabletEvent) TangentialPressure() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent18tangentialPressureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:282
// index:0
// Public inline
// qreal rotation()
func (this *QTabletEvent) Rotation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8rotationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:283
// index:0
// Public inline
// int xTilt()
func (this *QTabletEvent) XTilt() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent5xTiltEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:284
// index:0
// Public inline
// int yTilt()
func (this *QTabletEvent) YTilt() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent5yTiltEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:285
// index:0
// Public
// Qt::MouseButton button()
func (this *QTabletEvent) Button() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:286
// index:0
// Public
// Qt::MouseButtons buttons()
func (this *QTabletEvent) Buttons() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7buttonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
