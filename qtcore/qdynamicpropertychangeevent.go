package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QDynamicPropertyChangeEvent struct {
	*QEvent
}

func (this *QDynamicPropertyChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QDynamicPropertyChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = NewQEventFromPointer(cthis)
}
func NewQDynamicPropertyChangeEventFromPointer(cthis unsafe.Pointer) *QDynamicPropertyChangeEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QDynamicPropertyChangeEvent{bcthis0}
}
func (*QDynamicPropertyChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QDynamicPropertyChangeEvent {
	return NewQDynamicPropertyChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:365
// index:0
// Public
// void QDynamicPropertyChangeEvent(const class QByteArray &)
func NewQDynamicPropertyChangeEvent(name *QByteArray) *QDynamicPropertyChangeEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDynamicPropertyChangeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:366
// index:0
// Public virtual
// void ~QDynamicPropertyChangeEvent()
func DeleteQDynamicPropertyChangeEvent(*QDynamicPropertyChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:368
// index:0
// Public inline
// QByteArray propertyName()
func (this *QDynamicPropertyChangeEvent) PropertyName() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QDynamicPropertyChangeEvent12propertyNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
