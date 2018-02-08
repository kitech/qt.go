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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QInputMethodEvent struct {
	*qtcore.QEvent
}

func (this *QInputMethodEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QInputMethodEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQInputMethodEventFromPointer(cthis unsafe.Pointer) *QInputMethodEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QInputMethodEvent{bcthis0}
}
func (*QInputMethodEvent) NewFromPointer(cthis unsafe.Pointer) *QInputMethodEvent {
	return NewQInputMethodEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:555
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QInputMethodEvent()
func NewQInputMethodEvent() *QInputMethodEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QInputMethodEventC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQInputMethodEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQInputMethodEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:557
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QInputMethodEvent()
func DeleteQInputMethodEvent(this *QInputMethodEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QInputMethodEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 56)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:559
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCommitString(const QString &, int, int)
func (this *QInputMethodEvent) SetCommitString(commitString string, replaceFrom int, replaceLength int) {
	var tmpArg0 = qtcore.NewQString_5(commitString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QInputMethodEvent15setCommitStringERK7QStringii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, replaceFrom, replaceLength)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:560
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QList<QInputMethodEvent::Attribute> & attributes()
func (this *QInputMethodEvent) Attributes() unsafe.Pointer /*555*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QInputMethodEvent10attributesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qevent.h:561
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString & preeditString()
func (this *QInputMethodEvent) PreeditString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QInputMethodEvent13preeditStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qevent.h:563
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString & commitString()
func (this *QInputMethodEvent) CommitString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QInputMethodEvent12commitStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qevent.h:564
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int replacementStart()
func (this *QInputMethodEvent) ReplacementStart() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QInputMethodEvent16replacementStartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:565
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int replacementLength()
func (this *QInputMethodEvent) ReplacementLength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QInputMethodEvent17replacementLengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QInputMethodEvent__AttributeType = int

const QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = 0
const QInputMethodEvent__Cursor QInputMethodEvent__AttributeType = 1
const QInputMethodEvent__Language QInputMethodEvent__AttributeType = 2
const QInputMethodEvent__Ruby QInputMethodEvent__AttributeType = 3
const QInputMethodEvent__Selection QInputMethodEvent__AttributeType = 4

//  body block end
