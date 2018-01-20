//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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

// /usr/include/qt/QtGui/qevent.h:250
// index:0
// void QTabletEvent(enum QEvent::Type, const class QPointF &, const class QPointF &, int, int, qreal, int, int, qreal, qreal, int, Qt::KeyboardModifiers, qint64)
func NewQTabletEvent(t int, pos unsafe.Pointer, globalPos unsafe.Pointer, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64) *QTabletEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTabletEventC2EN6QEvent4TypeERK7QPointFS4_iidiiddi6QFlagsIN2Qt16KeyboardModifierEEx", ffiqt.FFI_TYPE_VOID, cthis, &t, pos, globalPos, &device, &pointerType, &pressure, &xTilt, &yTilt, &tangentialPressure, &rotation, &z, &keyState, &uniqueID)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabletEventFromPointer(cthis)
	return gothis
}
func NewQTabletEventFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTabletEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:254
// index:1
// void QTabletEvent(enum QEvent::Type, const class QPointF &, const class QPointF &, int, int, qreal, int, int, qreal, qreal, int, Qt::KeyboardModifiers, qint64, Qt::MouseButton, Qt::MouseButtons)
func NewQTabletEvent_1(t int, pos unsafe.Pointer, globalPos unsafe.Pointer, device int, pointerType int, pressure float64, xTilt int, yTilt int, tangentialPressure float64, rotation float64, z int, keyState int, uniqueID int64, button int, buttons int) *QTabletEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTabletEventC2EN6QEvent4TypeERK7QPointFS4_iidiiddi6QFlagsIN2Qt16KeyboardModifierEExNS6_11MouseButtonES5_IS9_E", ffiqt.FFI_TYPE_VOID, cthis, &t, pos, globalPos, &device, &pointerType, &pressure, &xTilt, &yTilt, &tangentialPressure, &rotation, &z, &keyState, &uniqueID, &button, &buttons)
	gopp.ErrPrint(err, rv)
	gothis := NewQTabletEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:259
// index:0
// virtual
// void ~QTabletEvent()
func DeleteQTabletEvent(*QTabletEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QTabletEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:261
// index:0
// inline
// QPoint pos()
func (this *QTabletEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:262
// index:0
// inline
// QPoint globalPos()
func (this *QTabletEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:267
// index:0
// inline
// const QPointF & posF()
func (this *QTabletEvent) PosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent4posFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:268
// index:0
// inline
// const QPointF & globalPosF()
func (this *QTabletEvent) GlobalPosF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent10globalPosFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:270
// index:0
// inline
// int x()
func (this *QTabletEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1xEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:271
// index:0
// inline
// int y()
func (this *QTabletEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1yEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:272
// index:0
// inline
// int globalX()
func (this *QTabletEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:273
// index:0
// inline
// int globalY()
func (this *QTabletEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:274
// index:0
// inline
// qreal hiResGlobalX()
func (this *QTabletEvent) HiResGlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalXEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:275
// index:0
// inline
// qreal hiResGlobalY()
func (this *QTabletEvent) HiResGlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent12hiResGlobalYEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:276
// index:0
// inline
// QTabletEvent::TabletDevice device()
func (this *QTabletEvent) Device() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent6deviceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:277
// index:0
// inline
// QTabletEvent::PointerType pointerType()
func (this *QTabletEvent) PointerType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent11pointerTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:278
// index:0
// inline
// qint64 uniqueId()
func (this *QTabletEvent) UniqueId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8uniqueIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:279
// index:0
// inline
// qreal pressure()
func (this *QTabletEvent) Pressure() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8pressureEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:280
// index:0
// inline
// int z()
func (this *QTabletEvent) Z() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent1zEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:281
// index:0
// inline
// qreal tangentialPressure()
func (this *QTabletEvent) TangentialPressure() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent18tangentialPressureEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:282
// index:0
// inline
// qreal rotation()
func (this *QTabletEvent) Rotation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent8rotationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:283
// index:0
// inline
// int xTilt()
func (this *QTabletEvent) XTilt() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent5xTiltEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:284
// index:0
// inline
// int yTilt()
func (this *QTabletEvent) YTilt() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent5yTiltEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:285
// index:0
// Qt::MouseButton button()
func (this *QTabletEvent) Button() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent6buttonEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:286
// index:0
// Qt::MouseButtons buttons()
func (this *QTabletEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QTabletEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
