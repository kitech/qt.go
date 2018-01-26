package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
	*qtrt.CObject
}

func (this *QAccessibleEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleEvent) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQAccessibleEventFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return &QAccessibleEvent{&qtrt.CObject{cthis}}
}
func (*QAccessibleEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return NewQAccessibleEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:668
// index:0
// Public inline
// void QAccessibleEvent(class QObject *, class QAccessible::Event)
func NewQAccessibleEvent(obj *qtcore.QObject /*777 QObject **/, typ int) *QAccessibleEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP7QObjectN11QAccessible5EventE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, typ)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:684
// index:1
// Public inline
// void QAccessibleEvent(class QAccessibleInterface *, class QAccessible::Event)
func NewQAccessibleEvent_1(iface *QAccessibleInterface /*777 QAccessibleInterface **/, typ int) *QAccessibleEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = iface.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP20QAccessibleInterfaceN11QAccessible5EventE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, typ)
	gopp.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:699
// index:0
// Public virtual
// void ~QAccessibleEvent()
func DeleteQAccessibleEvent(*QAccessibleEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:701
// index:0
// Public inline
// QAccessible::Event type()
func (this *QAccessibleEvent) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:702
// index:0
// Public inline
// QObject * object()
func (this *QAccessibleEvent) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:703
// index:0
// Public
// QAccessible::Id uniqueId()
func (this *QAccessibleEvent) UniqueId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent8uniqueIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:705
// index:0
// Public inline
// void setChild(int)
func (this *QAccessibleEvent) SetChild(chld int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAccessibleEvent8setChildEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), chld)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:706
// index:0
// Public inline
// int child()
func (this *QAccessibleEvent) Child() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent5childEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qaccessible.h:708
// index:0
// Public virtual
// QAccessibleInterface * accessibleInterface()
func (this *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAccessibleEvent19accessibleInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
