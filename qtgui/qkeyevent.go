package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QKeyEvent struct {
	*QInputEvent
}
type QKeyEvent_ITF interface {
	QInputEvent_ITF
	QKeyEvent_PTR() *QKeyEvent
}

func (ptr *QKeyEvent) QKeyEvent_PTR() *QKeyEvent { return ptr }

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

// /usr/include/qt/QtGui/qevent.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, const QString &, _Bool, ushort)
func NewQKeyEvent(type_ int, key int, modifiers int, text string, autorep bool, count uint16) *QKeyEvent {
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEERK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, convArg3, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, const QString &, _Bool, ushort)
func NewQKeyEvent__(type_ int, key int, modifiers int) *QKeyEvent {
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = qtcore.NewQString()
	// arg: 4, bool=Bool, =Invalid,
	autorep := false
	// arg: 5, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEERK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, convArg3, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, const QString &, _Bool, ushort)
func NewQKeyEvent__1(type_ int, key int, modifiers int, text string) *QKeyEvent {
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, bool=Bool, =Invalid,
	autorep := false
	// arg: 5, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEERK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, convArg3, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, const QString &, _Bool, ushort)
func NewQKeyEvent__2(type_ int, key int, modifiers int, text string, autorep bool) *QKeyEvent {
	var tmpArg3 = qtcore.NewQString_5(text)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 5, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEERK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, convArg3, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:340
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, quint32, quint32, quint32, const QString &, _Bool, ushort)
func NewQKeyEvent_1(type_ int, key int, modifiers int, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	var tmpArg6 = qtcore.NewQString_5(text)
	var convArg6 = tmpArg6.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEEjjjRK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, convArg6, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:340
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, quint32, quint32, quint32, const QString &, _Bool, ushort)
func NewQKeyEvent_1_(type_ int, key int, modifiers int, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint) *QKeyEvent {
	// arg: 6, const QString &=LValueReference, QString=Record,
	var convArg6 = qtcore.NewQString()
	// arg: 7, bool=Bool, =Invalid,
	autorep := false
	// arg: 8, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEEjjjRK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, convArg6, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:340
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, quint32, quint32, quint32, const QString &, _Bool, ushort)
func NewQKeyEvent_1_1(type_ int, key int, modifiers int, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string) *QKeyEvent {
	var tmpArg6 = qtcore.NewQString_5(text)
	var convArg6 = tmpArg6.GetCthis()
	// arg: 7, bool=Bool, =Invalid,
	autorep := false
	// arg: 8, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEEjjjRK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, convArg6, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:340
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeyEvent(enum QEvent::Type, int, Qt::KeyboardModifiers, quint32, quint32, quint32, const QString &, _Bool, ushort)
func NewQKeyEvent_1_2(type_ int, key int, modifiers int, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool) *QKeyEvent {
	var tmpArg6 = qtcore.NewQString_5(text)
	var convArg6 = tmpArg6.GetCthis()
	// arg: 8, ushort=Typedef, ushort=Typedef, unsigned short
	count := 1
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventC2EN6QEvent4TypeEi6QFlagsIN2Qt16KeyboardModifierEEjjjRK7QStringbt", qtrt.FFI_TYPE_POINTER, type_, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, convArg6, autorep, count)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeyEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQKeyEvent)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:343
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QKeyEvent()
func DeleteQKeyEvent(this *QKeyEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QKeyEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 64)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qevent.h:345
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int key() const
func (this *QKeyEvent) Key() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:347
// index:0
// Public Visibility=Default Availability=Available
// [1] bool matches(QKeySequence::StandardKey) const
func (this *QKeyEvent) Matches(key int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent7matchesEN12QKeySequence11StandardKeyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:349
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifiers() const
func (this *QKeyEvent) Modifiers() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent9modifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qevent.h:350
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString text() const
func (this *QKeyEvent) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qevent.h:351
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAutoRepeat() const
func (this *QKeyEvent) IsAutoRepeat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent12isAutoRepeatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:352
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int count() const
func (this *QKeyEvent) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qevent.h:354
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeScanCode() const
func (this *QKeyEvent) NativeScanCode() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent14nativeScanCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeVirtualKey() const
func (this *QKeyEvent) NativeVirtualKey() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent16nativeVirtualKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qevent.h:356
// index:0
// Public inline Visibility=Default Availability=Available
// [4] quint32 nativeModifiers() const
func (this *QKeyEvent) NativeModifiers() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QKeyEvent15nativeModifiersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
