package qtwidgets

// /usr/include/qt/QtWidgets/qkeysequenceedit.h
// #include <qkeysequenceedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QKeySequenceEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QKeySequenceEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:58
// index:0
// Public
// void QKeySequenceEdit(class QWidget *)
func NewQKeySequenceEdit(parent *QWidget /*777 QWidget **/) *QKeySequenceEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:59
// index:1
// Public
// void QKeySequenceEdit(const class QKeySequence &, class QWidget *)
func NewQKeySequenceEdit_1(keySequence *qtgui.QKeySequence, parent *QWidget /*777 QWidget **/) *QKeySequenceEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = keySequence.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditC2ERK12QKeySequenceP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeySequenceEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:60
// index:0
// Public virtual
// void ~QKeySequenceEdit()
func DeleteQKeySequenceEdit(*QKeySequenceEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:62
// index:0
// Public
// QKeySequence keySequence()
func (this *QKeySequenceEdit) KeySequence() *qtgui.QKeySequence /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QKeySequenceEdit11keySequenceEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQKeySequenceFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:65
// index:0
// Public
// void setKeySequence(const class QKeySequence &)
func (this *QKeySequenceEdit) SetKeySequence(keySequence *qtgui.QKeySequence) {
	var convArg0 = keySequence.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit14setKeySequenceERK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:66
// index:0
// Public
// void clear()
func (this *QKeySequenceEdit) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:69
// index:0
// Public
// void editingFinished()
func (this *QKeySequenceEdit) EditingFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit15editingFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:70
// index:0
// Public
// void keySequenceChanged(const class QKeySequence &)
func (this *QKeySequenceEdit) KeySequenceChanged(keySequence *qtgui.QKeySequence) {
	var convArg0 = keySequence.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit18keySequenceChangedERK12QKeySequence", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:75
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QKeySequenceEdit) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:76
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:77
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QKeySequenceEdit) KeyReleaseEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeysequenceedit.h:78
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QKeySequenceEdit) TimerEvent(arg0 *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
