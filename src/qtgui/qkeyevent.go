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
	cthis unsafe.Pointer
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:347
// index:0
// bool matches(class QKeySequence::StandardKey)
func (this *QKeyEvent) Matches(key int) {
	// 0: (, QKeySequence::StandardKey key), (&key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent7matchesEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_VOID, this.cthis, &key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:349
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QKeyEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:350
// index:0
// inline
// QString text()
func (this *QKeyEvent) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:351
// index:0
// inline
// bool isAutoRepeat()
func (this *QKeyEvent) IsAutoRepeat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent12isAutoRepeatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:352
// index:0
// inline
// int count()
func (this *QKeyEvent) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:354
// index:0
// inline
// quint32 nativeScanCode()
func (this *QKeyEvent) NativeScanCode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent14nativeScanCodeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:355
// index:0
// inline
// quint32 nativeVirtualKey()
func (this *QKeyEvent) NativeVirtualKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent16nativeVirtualKeyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:356
// index:0
// inline
// quint32 nativeModifiers()
func (this *QKeyEvent) NativeModifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent15nativeModifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
