package qtwidgets

// /usr/include/qt/QtWidgets/qmouseeventtransition.h
// #include <qmouseeventtransition.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
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

// void onTransition(class QEvent *)
func (this *QMouseEventTransition) InheritOnTransition(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool eventTest(class QEvent *)
func (this *QMouseEventTransition) InheritEventTest(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

type QMouseEventTransition struct {
	*qtcore.QEventTransition
}
type QMouseEventTransition_ITF interface {
	qtcore.QEventTransition_ITF
	QMouseEventTransition_PTR() *QMouseEventTransition
}

func (ptr *QMouseEventTransition) QMouseEventTransition_PTR() *QMouseEventTransition { return ptr }

func (this *QMouseEventTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QEventTransition.GetCthis()
	}
}
func (this *QMouseEventTransition) SetCthis(cthis unsafe.Pointer) {
	this.QEventTransition = qtcore.NewQEventTransitionFromPointer(cthis)
}
func NewQMouseEventTransitionFromPointer(cthis unsafe.Pointer) *QMouseEventTransition {
	bcthis0 := qtcore.NewQEventTransitionFromPointer(cthis)
	return &QMouseEventTransition{bcthis0}
}
func (*QMouseEventTransition) NewFromPointer(cthis unsafe.Pointer) *QMouseEventTransition {
	return NewQMouseEventTransitionFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMouseEventTransition) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMouseEventTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMouseEventTransition(QState *)
func NewQMouseEventTransition(sourceState qtcore.QState_ITF /*777 QState **/) *QMouseEventTransition {
	var convArg0 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg0 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMouseEventTransition")
	return gothis
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMouseEventTransition(QObject *, QEvent::Type, Qt::MouseButton, QState *)
func NewQMouseEventTransition_1(object qtcore.QObject_ITF /*777 QObject **/, type_ int, button int, sourceState qtcore.QState_ITF /*777 QState **/) *QMouseEventTransition {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg3 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransitionC2EP7QObjectN6QEvent4TypeEN2Qt11MouseButtonEP6QState", qtrt.FFI_TYPE_POINTER, convArg0, type_, button, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMouseEventTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMouseEventTransition")
	return gothis
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMouseEventTransition()
func DeleteQMouseEventTransition(this *QMouseEventTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButton button()
func (this *QMouseEventTransition) Button() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMouseEventTransition6buttonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setButton(Qt::MouseButton)
func (this *QMouseEventTransition) SetButton(button int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransition9setButtonEN2Qt11MouseButtonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), button)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::KeyboardModifiers modifierMask()
func (this *QMouseEventTransition) ModifierMask() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMouseEventTransition12modifierMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModifierMask(Qt::KeyboardModifiers)
func (this *QMouseEventTransition) SetModifierMask(modifiers int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransition15setModifierMaskE6QFlagsIN2Qt16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modifiers)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath hitTestPath()
func (this *QMouseEventTransition) HitTestPath() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QMouseEventTransition11hitTestPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHitTestPath(const QPainterPath &)
func (this *QMouseEventTransition) SetHitTestPath(path qtgui.QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransition14setHitTestPathERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QMouseEventTransition) OnTransition(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmouseeventtransition.h:74
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QMouseEventTransition) EventTest(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QMouseEventTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
