//  header block begin
// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>
package qtcore

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
type QChildEvent struct {
	*QEvent
}

func (this *QChildEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}

// /usr/include/qt/QtCore/qcoreevent.h:352
// index:0
// void QChildEvent(enum QEvent::Type, class QObject *)
func NewQChildEvent(type_ int, child unsafe.Pointer) *QChildEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QChildEventC2EN6QEvent4TypeEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &type_, child)
	gopp.ErrPrint(err, rv)
	gothis := NewQChildEventFromPointer(cthis)
	return gothis
}
func NewQChildEventFromPointer(cthis unsafe.Pointer) *QChildEvent {
	bcthis0 := NewQEventFromPointer(cthis)
	return &QChildEvent{bcthis0}
}

// /usr/include/qt/QtCore/qcoreevent.h:353
// index:0
// virtual
// void ~QChildEvent()
func DeleteQChildEvent(*QChildEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QChildEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:354
// index:0
// inline
// QObject * child()
func (this *QChildEvent) Child() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QChildEvent5childEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:355
// index:0
// inline
// bool added()
func (this *QChildEvent) Added() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QChildEvent5addedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:356
// index:0
// inline
// bool polished()
func (this *QChildEvent) Polished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QChildEvent8polishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:357
// index:0
// inline
// bool removed()
func (this *QChildEvent) Removed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QChildEvent7removedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
