package qtwidgets

// /usr/include/qt/QtWidgets/qkeysequenceedit.h
// #include <qkeysequenceedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QKeySequenceEdit) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void timerEvent(class QTimerEvent *)
func (this *QKeySequenceEdit) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

type QKeySequenceEdit struct {
	*QWidget
}

func (this *QKeySequenceEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QKeySequenceEdit) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQKeySequenceEditFromPointer(cthis unsafe.Pointer) *QKeySequenceEdit {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QKeySequenceEdit{bcthis0}
}
func (*QKeySequenceEdit) NewFromPointer(cthis unsafe.Pointer) *QKeySequenceEdit {
	return NewQKeySequenceEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QKeySequenceEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QKeySequenceEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(QWidget *)
func NewQKeySequenceEdit(parent *QWidget /*777 QWidget **/) *QKeySequenceEdit {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(const QKeySequence &, QWidget *)
func NewQKeySequenceEdit_1(keySequence *qtgui.QKeySequence, parent *QWidget /*777 QWidget **/) *QKeySequenceEdit {
	var convArg0 = keySequence.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QKeySequenceEdit()
func DeleteQKeySequenceEdit(this *QKeySequenceEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence keySequence()
func (this *QKeySequenceEdit) KeySequence() *qtgui.QKeySequence /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QKeySequenceEdit11keySequenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeySequence(const QKeySequence &)
func (this *QKeySequenceEdit) SetKeySequence(keySequence *qtgui.QKeySequence) {
	var convArg0 = keySequence.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QKeySequenceEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()
func (this *QKeySequenceEdit) EditingFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit15editingFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void keySequenceChanged(const QKeySequence &)
func (this *QKeySequenceEdit) KeySequenceChanged(keySequence *qtgui.QKeySequence) {
	var convArg0 = keySequence.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit18keySequenceChangedERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:75
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QKeySequenceEdit) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QKeySequenceEdit) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QKeySequenceEdit) KeyReleaseEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QKeySequenceEdit) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
