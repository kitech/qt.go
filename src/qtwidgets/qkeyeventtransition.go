//  header block begin
// /usr/include/qt/QtWidgets/qkeyeventtransition.h
// #include <qkeyeventtransition.h>
// #include <QtWidgets>
package qtwidgets

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
type QKeyEventTransition struct {
	*qtcore.QEventTransition
}

func (this *QKeyEventTransition) GetCthis() unsafe.Pointer {
	return this.QEventTransition.GetCthis()
}
func NewQKeyEventTransitionFromPointer(cthis unsafe.Pointer) *QKeyEventTransition {
	bcthis0 := qtcore.NewQEventTransitionFromPointer(cthis)
	return &QKeyEventTransition{bcthis0}
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QKeyEventTransition) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:57
// index:0
// Public
// void QKeyEventTransition(class QState *)
func NewQKeyEventTransition(sourceState unsafe.Pointer) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQKeyEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:58
// index:1
// Public
// void QKeyEventTransition(class QObject *, class QEvent::Type, int, class QState *)
func NewQKeyEventTransition_1(object unsafe.Pointer, type_ int, key int, sourceState unsafe.Pointer) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP7QObjectN6QEvent4TypeEiP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, &key, sourceState)
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
func (this *QKeyEventTransition) Key() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:63
// index:0
// Public
// void setKey(int)
func (this *QKeyEventTransition) SetKey(key int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition6setKeyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:65
// index:0
// Public
// Qt::KeyboardModifiers modifierMask()
func (this *QKeyEventTransition) ModifierMask() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition12modifierMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:69
// index:0
// Protected virtual
// void onTransition(class QEvent *)
func (this *QKeyEventTransition) OnTransition(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:70
// index:0
// Protected virtual
// bool eventTest(class QEvent *)
func (this *QKeyEventTransition) EventTest(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
