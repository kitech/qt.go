//  header block begin
// /usr/include/qt/QtWidgets/qmouseeventtransition.h
// #include <qmouseeventtransition.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMouseEventTransition) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:58
// index:0
// void QMouseEventTransition(class QState *)
func NewQMouseEventTransition(sourceState unsafe.Pointer) *QMouseEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	return &QMouseEventTransition{cthis}
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:59
// index:1
// void QMouseEventTransition(class QObject *, class QEvent::Type, Qt::MouseButton, class QState *)
func NewQMouseEventTransition_1(object unsafe.Pointer, type_ int, button int, sourceState unsafe.Pointer) *QMouseEventTransition {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP7QObjectN6QEvent4TypeEN2Qt11MouseButtonEP6QState", ffiqt.FFI_TYPE_VOID, cthis, object, &type_, &button, sourceState)
	gopp.ErrPrint(err, rv)
	return &QMouseEventTransition{cthis}
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:61
// index:0
// virtual
// void ~QMouseEventTransition()
func DeleteQMouseEventTransition(*QMouseEventTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:63
// index:0
// Qt::MouseButton button()
func (this *QMouseEventTransition) Button() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition6buttonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:64
// index:0
// void setButton(Qt::MouseButton)
func (this *QMouseEventTransition) SetButton(button int) {
	// 0: (, Qt::MouseButton button), (&button)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition9setButtonEN2Qt11MouseButtonE", ffiqt.FFI_TYPE_VOID, this.cthis, &button)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:66
// index:0
// Qt::KeyboardModifiers modifierMask()
func (this *QMouseEventTransition) ModifierMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition12modifierMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:69
// index:0
// QPainterPath hitTestPath()
func (this *QMouseEventTransition) HitTestPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QMouseEventTransition11hitTestPathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:70
// index:0
// void setHitTestPath(const class QPainterPath &)
func (this *QMouseEventTransition) SetHitTestPath(path unsafe.Pointer) {
	// 0: (, const QPainterPath & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

//  body block end
