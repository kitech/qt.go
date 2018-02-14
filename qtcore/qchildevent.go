package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QChildEvent struct {
	*QEvent
}
type QChildEvent_ITF interface {
	QEvent_ITF
	QChildEvent_PTR() *QChildEvent
}

func (ptr *QChildEvent) QChildEvent_PTR() *QChildEvent { return ptr }

func (this *QChildEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEvent.GetCthis()
	}
}
func (this *QChildEvent) SetCthis(cthis unsafe.Pointer) {
	this.QEvent = NewQEventFromPointer(cthis)
}
func NewQChildEventFromPointer(cthis unsafe.Pointer) *QChildEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QChildEvent{bcthis0}
}
func (*QChildEvent) NewFromPointer(cthis unsafe.Pointer) *QChildEvent {
	return NewQChildEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:352
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QChildEvent(enum QEvent::Type, QObject *)
func NewQChildEvent(type_ int, child QObject_ITF /*777 QObject **/) *QChildEvent {
	var convArg1 = child.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QChildEventC2EN6QEvent4TypeEP7QObject", qtrt.FFI_TYPE_POINTER, type_, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQChildEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQChildEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:353
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QChildEvent()
func DeleteQChildEvent(this *QChildEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QChildEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:354
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * child()
func (this *QChildEvent) Child() *QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QChildEvent5childEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qcoreevent.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool added()
func (this *QChildEvent) Added() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QChildEvent5addedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:356
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool polished()
func (this *QChildEvent) Polished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QChildEvent8polishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:357
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removed()
func (this *QChildEvent) Removed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QChildEvent7removedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		qtrt.KeepMe()
	}
}

//  keep block end
