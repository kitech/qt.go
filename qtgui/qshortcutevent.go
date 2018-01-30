package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QShortcutEvent struct {
	*qtcore.QEvent
}

func (this *QShortcutEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QShortcutEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = qtcore.NewQEventFromPointer(cthis)
}
func NewQShortcutEventFromPointer(cthis unsafe.Pointer) *QShortcutEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QShortcutEvent{bcthis0}
}
func (*QShortcutEvent) NewFromPointer(cthis unsafe.Pointer) *QShortcutEvent {
	return NewQShortcutEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:767
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QShortcutEvent(const QKeySequence &, int, _Bool)
func NewQShortcutEvent(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventC2ERK12QKeySequenceib", ffiqt.FFI_TYPE_POINTER, convArg0, id, ambiguous)
	gopp.ErrPrint(err, rv)
	gothis := NewQShortcutEventFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:768
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QShortcutEvent()
func DeleteQShortcutEvent(*QShortcutEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QShortcutEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:770
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QKeySequence & key()
func (this *QShortcutEvent) Key() *QKeySequence {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:771
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int shortcutId()
func (this *QShortcutEvent) ShortcutId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent10shortcutIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:772
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAmbiguous()
func (this *QShortcutEvent) IsAmbiguous() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QShortcutEvent11isAmbiguousEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
