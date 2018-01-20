//  header block begin
// /usr/include/qt/QtCore/qabstracttransition.h
// #include <qabstracttransition.h>
// #include <QtCore>
package qtcore

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
}

//  ext block end

//  body block begin
type QAbstractTransition struct {
	*QObject
}

func (this *QAbstractTransition) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQAbstractTransitionFromPointer(cthis unsafe.Pointer) *QAbstractTransition {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractTransition{bcthis0}
}

// /usr/include/qt/QtCore/qabstracttransition.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractTransition) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:75
// index:0
// Public
// void QAbstractTransition(class QState *)
func NewQAbstractTransition(sourceState unsafe.Pointer) *QAbstractTransition {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransitionC1EP6QState", ffiqt.FFI_TYPE_VOID, cthis, sourceState)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractTransitionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstracttransition.h:76
// index:0
// Public virtual
// void ~QAbstractTransition()
func DeleteQAbstractTransition(*QAbstractTransition) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransitionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:78
// index:0
// Public
// QState * sourceState()
func (this *QAbstractTransition) SourceState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition11sourceStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:79
// index:0
// Public
// QAbstractState * targetState()
func (this *QAbstractTransition) TargetState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition11targetStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:80
// index:0
// Public
// void setTargetState(class QAbstractState *)
func (this *QAbstractTransition) SetTargetState(target unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition14setTargetStateEP14QAbstractState", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:81
// index:0
// Public
// QList<QAbstractState *> targetStates()
func (this *QAbstractTransition) TargetStates() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition12targetStatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:84
// index:0
// Public
// QAbstractTransition::TransitionType transitionType()
func (this *QAbstractTransition) TransitionType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition14transitionTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:85
// index:0
// Public
// void setTransitionType(enum QAbstractTransition::TransitionType)
func (this *QAbstractTransition) SetTransitionType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition17setTransitionTypeENS_14TransitionTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:87
// index:0
// Public
// QStateMachine * machine()
func (this *QAbstractTransition) Machine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition7machineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:90
// index:0
// Public
// void addAnimation(class QAbstractAnimation *)
func (this *QAbstractTransition) AddAnimation(animation unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition12addAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), animation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:91
// index:0
// Public
// void removeAnimation(class QAbstractAnimation *)
func (this *QAbstractTransition) RemoveAnimation(animation unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), animation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:92
// index:0
// Public
// QList<QAbstractAnimation *> animations()
func (this *QAbstractTransition) Animations() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractTransition10animationsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:101
// index:0
// Protected pure virtual
// bool eventTest(class QEvent *)
func (this *QAbstractTransition) EventTest(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition9eventTestEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstracttransition.h:103
// index:0
// Protected pure virtual
// void onTransition(class QEvent *)
func (this *QAbstractTransition) OnTransition(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition12onTransitionEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracttransition.h:105
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractTransition) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractTransition5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
