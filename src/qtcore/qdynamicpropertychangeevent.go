//  header block begin
// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>
package qtcore

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtCore/qcoreevent.h:365
// index:0
// void QDynamicPropertyChangeEvent(const class QByteArray &)
func NewQDynamicPropertyChangeEvent(name unsafe.Pointer) *QDynamicPropertyChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	gothis := NewQDynamicPropertyChangeEventFromPointer(cthis)
	return gothis
}
func NewQDynamicPropertyChangeEventFromPointer(cthis unsafe.Pointer) *QDynamicPropertyChangeEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QDynamicPropertyChangeEvent{bcthis0}
}

// /usr/include/qt/QtCore/qcoreevent.h:366
// index:0
// virtual
// void ~QDynamicPropertyChangeEvent()
func DeleteQDynamicPropertyChangeEvent(*QDynamicPropertyChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN27QDynamicPropertyChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:368
// index:0
// inline
// QByteArray propertyName()
func (this *QDynamicPropertyChangeEvent) PropertyName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK27QDynamicPropertyChangeEvent12propertyNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
