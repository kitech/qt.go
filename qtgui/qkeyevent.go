package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

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

type QKeyEvent struct {
	*QInputEvent
}

func (this *QKeyEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QInputEvent.GetCthis()
	}
}
func (this *QKeyEvent) SetCthis(cthis unsafe.Pointer) {
	this.QInputEvent = NewQInputEventFromPointer(cthis)
}
func NewQKeyEventFromPointer(cthis unsafe.Pointer) *QKeyEvent {
	bcthis0 := NewQInputEventFromPointer(cthis)
	return &QKeyEvent{bcthis0}
}
func (*QKeyEvent) NewFromPointer(cthis unsafe.Pointer) *QKeyEvent {
	return NewQKeyEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:343
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QKeyEvent()
func DeleteQKeyEvent(this *QKeyEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QKeyEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:345
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int key()
func (this *QKeyEvent) Key() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:347
// index:0
// Public Visibility=Default Availability=Available
// [1] bool matches(QKeySequence::StandardKey)
func (this *QKeyEvent) Matches(key int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent7matchesEN12QKeySequence11StandardKeyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:349
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers()
func (this *QKeyEvent) Modifiers() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent9modifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:350
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text()
func (this *QKeyEvent) Text() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:351
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAutoRepeat()
func (this *QKeyEvent) IsAutoRepeat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent12isAutoRepeatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:352
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count()
func (this *QKeyEvent) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:354
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeScanCode()
func (this *QKeyEvent) NativeScanCode() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent14nativeScanCodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeVirtualKey()
func (this *QKeyEvent) NativeVirtualKey() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent16nativeVirtualKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:356
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeModifiers()
func (this *QKeyEvent) NativeModifiers() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QKeyEvent15nativeModifiersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

//  body block end
