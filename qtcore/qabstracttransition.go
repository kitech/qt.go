package qtcore

// /usr/include/qt/QtCore/qabstracttransition.h
// #include <qabstracttransition.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

// bool eventTest(class QEvent *)
func (this *QAbstractTransition) InheritEventTest(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventTest", f)
}

// void onTransition(class QEvent *)
func (this *QAbstractTransition) InheritOnTransition(f func(event *QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "onTransition", f)
}

// bool event(class QEvent *)
func (this *QAbstractTransition) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QAbstractTransition struct {
	*QObject
}
type QAbstractTransition_ITF interface {
	QObject_ITF
	QAbstractTransition_PTR() *QAbstractTransition
}

func (ptr *QAbstractTransition) QAbstractTransition_PTR() *QAbstractTransition { return ptr }

func (this *QAbstractTransition) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractTransition) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractTransitionFromPointer(cthis unsafe.Pointer) *QAbstractTransition {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractTransition{bcthis0}
}
func (*QAbstractTransition) NewFromPointer(cthis unsafe.Pointer) *QAbstractTransition {
	return NewQAbstractTransitionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstracttransition.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QAbstractTransition) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTransition10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracttransition.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTransition(QState *)
func NewQAbstractTransition(sourceState QState_ITF /*777 QState **/) *QAbstractTransition {
	var convArg0 unsafe.Pointer
	if sourceState != nil && sourceState.QState_PTR() != nil {
		convArg0 = sourceState.QState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractTransition")
	return gothis
}

// /usr/include/qt/QtCore/qabstracttransition.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTransition(QState *)
func NewQAbstractTransition__() *QAbstractTransition {
	// arg: 0, QState *=Pointer, QState=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransitionC2EP6QState", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractTransitionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractTransition")
	return gothis
}

// /usr/include/qt/QtCore/qabstracttransition.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractTransition()
func DeleteQAbstractTransition(this *QAbstractTransition) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransitionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstracttransition.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QState * sourceState() const
func (this *QAbstractTransition) SourceState() *QState /*777 QState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTransition11sourceStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracttransition.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractState * targetState() const
func (this *QAbstractTransition) TargetState() *QAbstractState /*777 QAbstractState **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTransition11targetStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractStateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracttransition.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTargetState(QAbstractState *)
func (this *QAbstractTransition) SetTargetState(target QAbstractState_ITF /*777 QAbstractState **/) {
	var convArg0 unsafe.Pointer
	if target != nil && target.QAbstractState_PTR() != nil {
		convArg0 = target.QAbstractState_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition14setTargetStateEP14QAbstractState", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractTransition::TransitionType transitionType() const
func (this *QAbstractTransition) TransitionType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTransition14transitionTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransitionType(enum QAbstractTransition::TransitionType)
func (this *QAbstractTransition) SetTransitionType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition17setTransitionTypeENS_14TransitionTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QStateMachine * machine() const
func (this *QAbstractTransition) Machine() *QStateMachine /*777 QStateMachine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTransition7machineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStateMachineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracttransition.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnimation(QAbstractAnimation *)
func (this *QAbstractTransition) AddAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAnimation(QAbstractAnimation *)
func (this *QAbstractTransition) RemoveAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:101
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [1] bool eventTest(QEvent *)
func (this *QAbstractTransition) EventTest(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition9eventTestEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracttransition.h:103
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [-2] void onTransition(QEvent *)
func (this *QAbstractTransition) OnTransition(event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition12onTransitionEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:105
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QAbstractTransition) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QAbstractTransition__TransitionType = int

const QAbstractTransition__ExternalTransition QAbstractTransition__TransitionType = 0
const QAbstractTransition__InternalTransition QAbstractTransition__TransitionType = 1

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
}

//  keep block end
