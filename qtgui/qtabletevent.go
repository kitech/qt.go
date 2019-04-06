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
// extern C begin: 24
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
type QTabletEvent struct {
	*QInputEvent
}
type QTabletEvent_ITF interface {
	QInputEvent_ITF
	QTabletEvent_PTR() *QTabletEvent
}

func (ptr *QTabletEvent) QTabletEvent_PTR() *QTabletEvent { return ptr }

func (this *QTabletEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QTabletEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQTabletEventFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QTabletEvent{bcthis0}
}
func (*QTabletEvent) NewFromPointer(cthis unsafe.Pointer) *QTabletEvent {
	return NewQTabletEventFromPointer(cthis)
}

/*


 */
type QTabletEvent__TabletDevice = int

//
const QTabletEvent__NoDevice QTabletEvent__TabletDevice = 0

//
const QTabletEvent__Puck QTabletEvent__TabletDevice = 1

//
const QTabletEvent__Stylus QTabletEvent__TabletDevice = 2

//
const QTabletEvent__Airbrush QTabletEvent__TabletDevice = 3

//
const QTabletEvent__FourDMouse QTabletEvent__TabletDevice = 4

//
const QTabletEvent__XFreeEraser QTabletEvent__TabletDevice = 5

//
const QTabletEvent__RotationStylus QTabletEvent__TabletDevice = 6

func (this *QTabletEvent) TabletDeviceItemName(val int) string {
	switch val {
	case QTabletEvent__NoDevice: // 0
		return "NoDevice"
	case QTabletEvent__Puck: // 1
		return "Puck"
	case QTabletEvent__Stylus: // 2
		return "Stylus"
	case QTabletEvent__Airbrush: // 3
		return "Airbrush"
	case QTabletEvent__FourDMouse: // 4
		return "FourDMouse"
	case QTabletEvent__XFreeEraser: // 5
		return "XFreeEraser"
	case QTabletEvent__RotationStylus: // 6
		return "RotationStylus"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTabletEvent_TabletDeviceItemName(val int) string {
	var nilthis *QTabletEvent
	return nilthis.TabletDeviceItemName(val)
}

/*


 */
type QTabletEvent__PointerType = int

//
const QTabletEvent__UnknownPointer QTabletEvent__PointerType = 0

//
const QTabletEvent__Pen QTabletEvent__PointerType = 1

//
const QTabletEvent__Cursor QTabletEvent__PointerType = 2

//
const QTabletEvent__Eraser QTabletEvent__PointerType = 3

func (this *QTabletEvent) PointerTypeItemName(val int) string {
	switch val {
	case QTabletEvent__UnknownPointer: // 0
		return "UnknownPointer"
	case QTabletEvent__Pen: // 1
		return "Pen"
	case QTabletEvent__Cursor: // 2
		return "Cursor"
	case QTabletEvent__Eraser: // 3
		return "Eraser"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTabletEvent_PointerTypeItemName(val int) string {
	var nilthis *QTabletEvent
	return nilthis.PointerTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10653() {
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
