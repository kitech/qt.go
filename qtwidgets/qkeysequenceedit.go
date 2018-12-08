package qtwidgets

// /usr/include/qt/QtWidgets/qkeysequenceedit.h
// #include <qkeysequenceedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QKeySequenceEdit) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QKeySequenceEdit) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QKeySequenceEdit) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void timerEvent(QTimerEvent *)
func (this *QKeySequenceEdit) InheritTimerEvent(f func(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

/*

 */
type QKeySequenceEdit struct {
	*QWidget
}
type QKeySequenceEdit_ITF interface {
	QWidget_ITF
	QKeySequenceEdit_PTR() *QKeySequenceEdit
}

func (ptr *QKeySequenceEdit) QKeySequenceEdit_PTR() *QKeySequenceEdit { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QKeySequenceEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QKeySequenceEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(QWidget *)

/*
Constructs a QKeySequenceEdit widget with the given parent.
*/
func (*QKeySequenceEdit) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QKeySequenceEdit {
	return NewQKeySequenceEdit(parent)
}
func NewQKeySequenceEdit(parent QWidget_ITF /*777 QWidget **/) *QKeySequenceEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QKeySequenceEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(QWidget *)

/*
Constructs a QKeySequenceEdit widget with the given parent.
*/
func (*QKeySequenceEdit) NewForInheritp() *QKeySequenceEdit {
	return NewQKeySequenceEditp()
}
func NewQKeySequenceEditp() *QKeySequenceEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QKeySequenceEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(const QKeySequence &, QWidget *)

/*
Constructs a QKeySequenceEdit widget with the given parent.
*/
func (*QKeySequenceEdit) NewForInherit1(keySequence qtgui.QKeySequence_ITF, parent QWidget_ITF /*777 QWidget **/) *QKeySequenceEdit {
	return NewQKeySequenceEdit1(keySequence, parent)
}
func NewQKeySequenceEdit1(keySequence qtgui.QKeySequence_ITF, parent QWidget_ITF /*777 QWidget **/) *QKeySequenceEdit {
	var convArg0 unsafe.Pointer
	if keySequence != nil && keySequence.QKeySequence_PTR() != nil {
		convArg0 = keySequence.QKeySequence_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QKeySequenceEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QKeySequenceEdit(const QKeySequence &, QWidget *)

/*
Constructs a QKeySequenceEdit widget with the given parent.
*/
func (*QKeySequenceEdit) NewForInherit1p(keySequence qtgui.QKeySequence_ITF) *QKeySequenceEdit {
	return NewQKeySequenceEdit1p(keySequence)
}
func NewQKeySequenceEdit1p(keySequence qtgui.QKeySequence_ITF) *QKeySequenceEdit {
	var convArg0 unsafe.Pointer
	if keySequence != nil && keySequence.QKeySequence_PTR() != nil {
		convArg0 = keySequence.QKeySequence_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QKeySequenceEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QKeySequenceEdit()

/*

 */
func DeleteQKeySequenceEdit(this *QKeySequenceEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] QKeySequence keySequence() const

/*

 */
func (this *QKeySequenceEdit) KeySequence() *qtgui.QKeySequence /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QKeySequenceEdit11keySequenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQKeySequence)
	return rv2
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeySequence(const QKeySequence &)

/*

 */
func (this *QKeySequenceEdit) SetKeySequence(keySequence qtgui.QKeySequence_ITF) {
	var convArg0 unsafe.Pointer
	if keySequence != nil && keySequence.QKeySequence_PTR() != nil {
		convArg0 = keySequence.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the current key sequence.
*/
func (this *QKeySequenceEdit) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void editingFinished()

/*
This signal is emitted when the user finishes entering the shortcut.

Note: there is a one second delay before releasing the last key and emitting this signal.
*/
func (this *QKeySequenceEdit) EditingFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit15editingFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void keySequenceChanged(const QKeySequence &)

/*

 */
func (this *QKeySequenceEdit) KeySequenceChanged(keySequence qtgui.QKeySequence_ITF) {
	var convArg0 unsafe.Pointer
	if keySequence != nil && keySequence.QKeySequence_PTR() != nil {
		convArg0 = keySequence.QKeySequence_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit18keySequenceChangedERK12QKeySequence", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:75
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QKeySequenceEdit) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyPressEvent().
*/
func (this *QKeySequenceEdit) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:77
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QWidget::keyReleaseEvent().
*/
func (this *QKeySequenceEdit) KeyReleaseEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)

/*
Reimplemented from QObject::timerEvent().
*/
func (this *QKeySequenceEdit) TimerEvent(arg0 qtcore.QTimerEvent_ITF /*777 QTimerEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QTimerEvent_PTR() != nil {
		convArg0 = arg0.QTimerEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
