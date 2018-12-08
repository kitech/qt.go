package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QAccessibleTextUpdateEvent struct {
	*QAccessibleTextCursorEvent
}
type QAccessibleTextUpdateEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextUpdateEvent_PTR() *QAccessibleTextUpdateEvent
}

func (ptr *QAccessibleTextUpdateEvent) QAccessibleTextUpdateEvent_PTR() *QAccessibleTextUpdateEvent {
	return ptr
}

func (this *QAccessibleTextUpdateEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleTextCursorEvent.GetCthis()
	}
}
func (this *QAccessibleTextUpdateEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleTextCursorEvent = NewQAccessibleTextCursorEventFromPointer(cthis)
}
func NewQAccessibleTextUpdateEventFromPointer(cthis unsafe.Pointer) *QAccessibleTextUpdateEvent {
	bcthis0 := NewQAccessibleTextCursorEventFromPointer(cthis)
	return &QAccessibleTextUpdateEvent{bcthis0}
}
func (*QAccessibleTextUpdateEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTextUpdateEvent {
	return NewQAccessibleTextUpdateEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:864
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextUpdateEvent(QObject *, int, const QString &, const QString &)

/*

 */
func (*QAccessibleTextUpdateEvent) NewForInherit(obj qtcore.QObject_ITF /*777 QObject **/, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return NewQAccessibleTextUpdateEvent(obj, position, oldText, text)
}
func NewQAccessibleTextUpdateEvent(obj qtcore.QObject_ITF /*777 QObject **/, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(oldText)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextUpdateEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextUpdateEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:870
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTextUpdateEvent(QAccessibleInterface *, int, const QString &, const QString &)

/*

 */
func (*QAccessibleTextUpdateEvent) NewForInherit1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return NewQAccessibleTextUpdateEvent1(iface, position, oldText, text)
}
func NewQAccessibleTextUpdateEvent1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(oldText)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_", qtrt.FFI_TYPE_POINTER, convArg0, position, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTextUpdateEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTextUpdateEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:877
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTextUpdateEvent()

/*

 */
func DeleteQAccessibleTextUpdateEvent(this *QAccessibleTextUpdateEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAccessibleTextUpdateEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:879
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textRemoved() const

/*

 */
func (this *QAccessibleTextUpdateEvent) TextRemoved() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent11textRemovedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:882
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString textInserted() const

/*

 */
func (this *QAccessibleTextUpdateEvent) TextInserted() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent12textInsertedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:885
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int changePosition() const

/*

 */
func (this *QAccessibleTextUpdateEvent) ChangePosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QAccessibleTextUpdateEvent14changePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
