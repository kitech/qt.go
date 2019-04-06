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
// extern C begin: 6
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
type QWheelEvent struct {
	*QInputEvent
}
type QWheelEvent_ITF interface {
	QInputEvent_ITF
	QWheelEvent_PTR() *QWheelEvent
}

func (ptr *QWheelEvent) QWheelEvent_PTR() *QWheelEvent { return ptr }

func (this *QWheelEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QWheelEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQWheelEventFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QWheelEvent{bcthis0}
}
func (*QWheelEvent) NewFromPointer(cthis unsafe.Pointer) *QWheelEvent {
	return NewQWheelEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & posF() const

/*

 */
func (this *QWheelEvent) PosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent4posFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [16] const QPointF & globalPosF() const

/*

 */
func (this *QWheelEvent) GlobalPosF() *qtcore.QPointF {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent10globalPosFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:220
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseButtons buttons() const

/*

 */
func (this *QWheelEvent) Buttons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent7buttonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ScrollPhase phase() const

/*

 */
func (this *QWheelEvent) Phase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent5phaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool inverted() const

/*

 */
func (this *QWheelEvent) Inverted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent8invertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::MouseEventSource source() const

/*

 */
func (this *QWheelEvent) Source() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QWheelEvent6sourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
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

func init_unused_10651() {
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
