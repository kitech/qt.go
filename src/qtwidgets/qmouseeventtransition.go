//  header block begin
// /usr/include/qt/QtWidgets/qmouseeventtransition.h
// #include <qmouseeventtransition.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 61
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
type QMouseEventTransition struct {
	*qtcore.QEventTransition
}

func (this *QMouseEventTransition) GetCthis() unsafe.Pointer {
	return this.QEventTransition.GetCthis()
}
func NewQMouseEventTransitionFromPointer(cthis unsafe.Pointer) *QMouseEventTransition {
	bcthis0 := qtcore.NewQEventTransitionFromPointer(cthis)
	return &QMouseEventTransition{bcthis0}
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QMouseEventTransition) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:58
// index:0
// Public
// void QMouseEventTransition(class QState *)
func NewQMouseEventTransition(sourceState unsafe.Pointer) *QMouseEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:59
// index:1
// Public
// void QMouseEventTransition(class QObject *, class QEvent::Type, Qt::MouseButton, class QState *)
func NewQMouseEventTransition_1(object unsafe.Pointer, type_ int, button int, sourceState unsafe.Pointer) *QMouseEventTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP7QObjectN6QEvent4TypeEN2Qt11MouseButtonEP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, &button, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQMouseEventTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:61
// index:0
// Public virtual
// void ~QMouseEventTransition()
func DeleteQMouseEventTransition(*QMouseEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:63
// index:0
// Public
// Qt::MouseButton button()
func (this *QMouseEventTransition) Button() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition6buttonEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:64
// index:0
// Public
// void setButton(Qt::MouseButton)
func (this *QMouseEventTransition) SetButton(button int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition9setButtonEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:66
// index:0
// Public
// Qt::KeyboardModifiers modifierMask()
func (this *QMouseEventTransition) ModifierMask() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition12modifierMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:69
// index:0
// Public
// QPainterPath hitTestPath()
func (this *QMouseEventTransition) HitTestPath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition11hitTestPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:70
// index:0
// Public
// void setHitTestPath(const class QPainterPath &)
func (this *QMouseEventTransition) SetHitTestPath(path *qtgui.QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:73
// index:0
// Protected virtual
// void onTransition(class QEvent *)
func (this *QMouseEventTransition) OnTransition(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:74
// index:0
// Protected virtual
// bool eventTest(class QEvent *)
func (this *QMouseEventTransition) EventTest(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
