//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QAccessibleEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qaccessible.h:668
// index:0
// inline
// void QAccessibleEvent(class QObject *, class QAccessible::Event)
func NewQAccessibleEvent(obj unsafe.Pointer, typ int) *QAccessibleEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP7QObjectN11QAccessible5EventE", ffiqt.FFI_TYPE_VOID, cthis, obj, &typ)
	gopp.ErrPrint(err, rv)
	return &QAccessibleEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:684
// index:1
// inline
// void QAccessibleEvent(class QAccessibleInterface *, class QAccessible::Event)
func NewQAccessibleEvent_1(iface unsafe.Pointer, typ int) *QAccessibleEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP20QAccessibleInterfaceN11QAccessible5EventE", ffiqt.FFI_TYPE_VOID, cthis, iface, &typ)
	gopp.ErrPrint(err, rv)
	return &QAccessibleEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:699
// index:0
// virtual
// void ~QAccessibleEvent()
func DeleteQAccessibleEvent(*QAccessibleEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:701
// index:0
// inline
// QAccessible::Event type()
func (this *QAccessibleEvent) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:702
// index:0
// inline
// QObject * object()
func (this *QAccessibleEvent) Object() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent6objectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:703
// index:0
// QAccessible::Id uniqueId()
func (this *QAccessibleEvent) UniqueId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent8uniqueIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:705
// index:0
// inline
// void setChild(int)
func (this *QAccessibleEvent) SetChild(chld int) {
	// 0: (, int chld), (&chld)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEvent8setChildEi", ffiqt.FFI_TYPE_VOID, this.cthis, &chld)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:706
// index:0
// inline
// int child()
func (this *QAccessibleEvent) Child() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent5childEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:708
// index:0
// virtual
// QAccessibleInterface * accessibleInterface()
func (this *QAccessibleEvent) AccessibleInterface() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent19accessibleInterfaceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
