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
// extern C begin: 2
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
type QHelpEvent struct {
	*qtcore.QEvent
}
type QHelpEvent_ITF interface {
	qtcore.QEvent_ITF
	QHelpEvent_PTR() *QHelpEvent
}

func (ptr *QHelpEvent) QHelpEvent_PTR() *QHelpEvent { return ptr }

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

// /usr/include/qt/QtGui/qevent.h:685
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QHelpEvent(QEvent::Type, const QPoint &, const QPoint &)

/*

 */
func (*QHelpEvent) NewForInherit(type_ int, pos qtcore.QPoint_ITF, globalPos qtcore.QPoint_ITF) *QHelpEvent {
	return NewQHelpEvent(type_, pos, globalPos)
}
func NewQHelpEvent(type_ int, pos qtcore.QPoint_ITF, globalPos qtcore.QPoint_ITF) *QHelpEvent {
	var convArg1 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg1 = pos.QPoint_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if globalPos != nil && globalPos.QPoint_PTR() != nil {
		convArg2 = globalPos.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QHelpEventC2EN6QEvent4TypeERK6QPointS4_", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQHelpEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQHelpEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:686
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QHelpEvent()

/*

 */
func DeleteQHelpEvent(this *QHelpEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QHelpEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:688
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x() const

/*

 */
func (this *QHelpEvent) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:689
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y() const

/*

 */
func (this *QHelpEvent) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:690
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalX() const

/*

 */
func (this *QHelpEvent) GlobalX() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent7globalXEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:691
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int globalY() const

/*

 */
func (this *QHelpEvent) GlobalY() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent7globalYEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:693
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & pos() const

/*

 */
func (this *QHelpEvent) Pos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:694
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QPoint & globalPos() const

/*

 */
func (this *QHelpEvent) GlobalPos() *qtcore.QPoint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QHelpEvent9globalPosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

//  body block end

//  keep block begin

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
