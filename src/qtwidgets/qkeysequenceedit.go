//  header block begin
// /usr/include/qt/QtWidgets/qkeysequenceedit.h
// #include <qkeysequenceedit.h>
// #include <QtWidgets>
package qtwidgets

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QKeySequenceEdit struct {
	*QWidget
}

func (this *QKeySequenceEdit) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QKeySequenceEdit) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QKeySequenceEdit10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:58
// index:0
// void QKeySequenceEdit(class QWidget *)
func NewQKeySequenceEdit(parent unsafe.Pointer) *QKeySequenceEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(cthis)
	return gothis
}
func NewQKeySequenceEditFromPointer(cthis unsafe.Pointer) *QKeySequenceEdit {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QKeySequenceEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:59
// index:1
// void QKeySequenceEdit(const class QKeySequence &, class QWidget *)
func NewQKeySequenceEdit_1(keySequence unsafe.Pointer, parent unsafe.Pointer) *QKeySequenceEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, keySequence, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:73
// index:2
// void QKeySequenceEdit(class QKeySequenceEditPrivate &, class QWidget *, Qt::WindowFlags)
func NewQKeySequenceEdit_2(d unsafe.Pointer, parent unsafe.Pointer, f int) *QKeySequenceEdit {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ER23QKeySequenceEditPrivateP7QWidget6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, d, parent, &f)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:60
// index:0
// virtual
// void ~QKeySequenceEdit()
func DeleteQKeySequenceEdit(*QKeySequenceEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:62
// index:0
// QKeySequence keySequence()
func (this *QKeySequenceEdit) KeySequence() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QKeySequenceEdit11keySequenceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:65
// index:0
// void setKeySequence(const class QKeySequence &)
func (this *QKeySequenceEdit) SetKeySequence(keySequence unsafe.Pointer) {
	// 0: (, keySequence const QKeySequence &), (keySequence)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), keySequence)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:66
// index:0
// void clear()
func (this *QKeySequenceEdit) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:69
// index:0
// void editingFinished()
func (this *QKeySequenceEdit) EditingFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit15editingFinishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:70
// index:0
// void keySequenceChanged(const class QKeySequence &)
func (this *QKeySequenceEdit) KeySequenceChanged(keySequence unsafe.Pointer) {
	// 0: (, keySequence const QKeySequence &), (keySequence)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit18keySequenceChangedERK12QKeySequence", ffiqt.FFI_TYPE_VOID, this.GetCthis(), keySequence)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:75
// index:0
// virtual
// bool event(class QEvent *)
func (this *QKeySequenceEdit) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:76
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:77
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) KeyReleaseEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:78
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QKeySequenceEdit) TimerEvent(arg0 unsafe.Pointer) {
	// 0: (, QTimerEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
