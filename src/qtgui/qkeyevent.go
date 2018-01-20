//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QKeyEvent struct {
	*QInputEvent
}

func (this *QKeyEvent) GetCthis() unsafe.Pointer {
	return this.QInputEvent.GetCthis()
}

// /usr/include/qt/QtGui/qevent.h:338
// index:0
// void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, const class QString &, _Bool, ushort)
func NewQKeyEvent(type_ int, key int, modifiers int, text unsafe.Pointer, autorep bool, count uint16) *QKeyEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEERK7QStringbt", ffiqt.FFI_TYPE_VOID, cthis, &type_, &key, &modifiers, text, &autorep, &count)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(cthis)
	return gothis
}
func NewQKeyEventFromPointer(cthis unsafe.Pointer) *QKeyEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QKeyEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:340
// index:1
// void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, quint32, quint32, quint32, const class QString &, _Bool, ushort)
func NewQKeyEvent_1(type_ int, key int, modifiers int, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text unsafe.Pointer, autorep bool, count uint16) *QKeyEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEEjjjRK7QStringbt", ffiqt.FFI_TYPE_VOID, cthis, &type_, &key, &modifiers, &nativeScanCode, &nativeVirtualKey, &nativeModifiers, text, &autorep, &count)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:343
// index:0
// virtual
// void ~QKeyEvent()
func DeleteQKeyEvent(*QKeyEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QKeyEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:345
// index:0
// inline
// int key()
func (this *QKeyEvent) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent3keyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:347
// index:0
// bool matches(class QKeySequence::StandardKey)
func (this *QKeyEvent) Matches(key int) {
	// 0: (, key QKeySequence::StandardKey), (&key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent7matchesEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:349
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QKeyEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:350
// index:0
// inline
// QString text()
func (this *QKeyEvent) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:351
// index:0
// inline
// bool isAutoRepeat()
func (this *QKeyEvent) IsAutoRepeat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent12isAutoRepeatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:352
// index:0
// inline
// int count()
func (this *QKeyEvent) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:354
// index:0
// inline
// quint32 nativeScanCode()
func (this *QKeyEvent) NativeScanCode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent14nativeScanCodeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:355
// index:0
// inline
// quint32 nativeVirtualKey()
func (this *QKeyEvent) NativeVirtualKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent16nativeVirtualKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:356
// index:0
// inline
// quint32 nativeModifiers()
func (this *QKeyEvent) NativeModifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent15nativeModifiersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
