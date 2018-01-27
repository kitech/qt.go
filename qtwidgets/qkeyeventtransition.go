package qtwidgets

// /usr/include/qt/QtWidgets/qkeyeventtransition.h
// #include <qkeyeventtransition.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type QKeyEventTransition struct {
	*qtcore.QEventTransition
}

func (this *QKeyEventTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEventTransition.GetCthis()
	}
}
func (this *QKeyEventTransition) SetCthis(cthis unsafe.Pointer) {
	this.QEventTransition = qtcore.NewQEventTransitionFromPointer(cthis)
}
func NewQKeyEventTransitionFromPointer(cthis unsafe.Pointer) *QKeyEventTransition {
	bcthis0 := qtcore.NewQEventTransitionFromPointer(cthis)
	return &QKeyEventTransition{bcthis0}
}
func (*QKeyEventTransition) NewFromPointer(cthis unsafe.Pointer) *QKeyEventTransition {
	return NewQKeyEventTransitionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QKeyEventTransition) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:57
// index:0
// Public
// void QKeyEventTransition(QState *)
func NewQKeyEventTransition(sourceState *qtcore.QState /*777 QState **/) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = sourceState.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeyEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:58
// index:1
// Public
// void QKeyEventTransition(QObject *, QEvent::Type, int, QState *)
func NewQKeyEventTransition_1(object *qtcore.QObject /*777 QObject **/, type_ int, key int, sourceState *qtcore.QState /*777 QState **/) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = object.GetCthis()
	var convArg3 = sourceState.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP7QObjectN6QEvent4TypeEiP6QState", ffiqt.FFI_TYPE_VOID, cthis, convArg0, type_, key, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeyEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:60
// index:0
// Public virtual
// void ~QKeyEventTransition()
func DeleteQKeyEventTransition(*QKeyEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:62
// index:0
// Public
// int key()
func (this *QKeyEventTransition) Key() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:63
// index:0
// Public
// void setKey(int)
func (this *QKeyEventTransition) SetKey(key int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition6setKeyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:65
// index:0
// Public
// Qt::KeyboardModifiers modifierMask()
func (this *QKeyEventTransition) ModifierMask() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition12modifierMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:69
// index:0
// Protected virtual
// void onTransition(QEvent *)
func (this *QKeyEventTransition) OnTransition(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:70
// index:0
// Protected virtual
// bool eventTest(QEvent *)
func (this *QKeyEventTransition) EventTest(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
