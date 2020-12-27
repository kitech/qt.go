package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 96
type QWheelEvent struct {
	*QInputEvent
}
type QWheelEvent_ITF interface {
	QInputEvent_ITF
	QWheelEvent_PTR() *QWheelEvent
}

func (ptr *QWheelEvent) QWheelEvent_PTR() *QWheelEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QWheelEventFromptr(cthis Voidptr) *QWheelEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QWheelEvent{bcthis0}
}
func (*QWheelEvent) Fromptr(cthis Voidptr) *QWheelEvent {
	return QWheelEventFromptr(cthis)
}

// /usr/include/qt/QtGui/qevent.h:243
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [16] QPointF position() const

/*
 */
func (this *QWheelEvent) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.Qtcc3(1239393956, "_ZNK11QWheelEvent8positionEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := qtcore.QPointFFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:244
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [16] QPointF globalPosition() const

/*
 */
func (this *QWheelEvent) GlobalPosition() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.Qtcc3(3973060595, "_ZNK11QWheelEvent14globalPositionEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(16)
	qtrt.Cmemcpy(cthis, rv.Addr(), 16)
	rv2 := qtcore.QPointFFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:249
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool inverted() const

/*
 */
func (this *QWheelEvent) Inverted() bool {
	rv, err := qtrt.Qtcc3(643606857, "_ZNK11QWheelEvent8invertedEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

/*


 */
type QWheelEvent__ = int

//
const QWheelEvent__DefaultDeltasPerStep QWheelEvent__ = 120

func (this *QWheelEvent) ItemName(val int) string {
	switch val {
	case QWheelEvent__DefaultDeltasPerStep: // 120
		return "DefaultDeltasPerStep"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWheelEvent_ItemName(val int) string {
	var nilthis *QWheelEvent
	return nilthis.ItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10075() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
