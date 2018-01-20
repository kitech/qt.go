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
func NewQKeyEventFromPointer(cthis unsafe.Pointer) *QKeyEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QKeyEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:343
// index:0
// Public virtual
// void ~QKeyEvent()
func DeleteQKeyEvent(*QKeyEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QKeyEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:345
// index:0
// Public inline
// int key()
func (this *QKeyEvent) Key() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:347
// index:0
// Public
// bool matches(class QKeySequence::StandardKey)
func (this *QKeyEvent) Matches(key int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent7matchesEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &key)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:349
// index:0
// Public
// Qt::KeyboardModifiers modifiers()
func (this *QKeyEvent) Modifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:350
// index:0
// Public inline
// QString text()
func (this *QKeyEvent) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:351
// index:0
// Public inline
// bool isAutoRepeat()
func (this *QKeyEvent) IsAutoRepeat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent12isAutoRepeatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:352
// index:0
// Public inline
// int count()
func (this *QKeyEvent) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:354
// index:0
// Public inline
// quint32 nativeScanCode()
func (this *QKeyEvent) NativeScanCode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent14nativeScanCodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:355
// index:0
// Public inline
// quint32 nativeVirtualKey()
func (this *QKeyEvent) NativeVirtualKey() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent16nativeVirtualKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:356
// index:0
// Public inline
// quint32 nativeModifiers()
func (this *QKeyEvent) NativeModifiers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent15nativeModifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
