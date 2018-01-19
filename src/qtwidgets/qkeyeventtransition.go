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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QKeyEventTransition) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:57
// index:0
// void QKeyEventTransition(class QState *)
func NewQKeyEventTransition(sourceState unsafe.Pointer) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	return &QKeyEventTransition{cthis}
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:58
// index:1
// void QKeyEventTransition(class QObject *, class QEvent::Type, int, class QState *)
func NewQKeyEventTransition_1(object unsafe.Pointer, type_ int, key int, sourceState unsafe.Pointer) *QKeyEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionC2EP7QObjectN6QEvent4TypeEiP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, &key, sourceState)
	gopp.ErrPrint(err, rv)
	return &QKeyEventTransition{cthis}
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:60
// index:0
// virtual
// void ~QKeyEventTransition()
func DeleteQKeyEventTransition(*QKeyEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:62
// index:0
// int key()
func (this *QKeyEventTransition) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:63
// index:0
// void setKey(int)
func (this *QKeyEventTransition) SetKey(key int) {
	// 0: (, int key), (&key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QKeyEventTransition6setKeyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qkeyeventtransition.h:65
// index:0
// Qt::KeyboardModifiers modifierMask()
func (this *QKeyEventTransition) ModifierMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QKeyEventTransition12modifierMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
