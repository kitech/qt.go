//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qstatemachine.h
// #include <qstatemachine.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qstatemachine.h:126
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAnimated() const

/*
Returns whether animations are enabled for this state machine.

Note: Getter function for property animated.
*/
func (this *QStateMachine) IsAnimated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStateMachine10isAnimatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstatemachine.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnimated(bool)

/*
Sets whether animations are enabled for this state machine.

Note: Setter function for property animated.

See also isAnimated().
*/
func (this *QStateMachine) SetAnimated(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11setAnimatedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addDefaultAnimation(QAbstractAnimation *)

/*
Adds a default animation to be considered for any transition.
*/
func (this *QStateMachine) AddDefaultAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine19addDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeDefaultAnimation(QAbstractAnimation *)

/*
Removes animation from the list of default animations.
*/
func (this *QStateMachine) RemoveDefaultAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine22removeDefaultAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstatemachine.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QStateMachine) EventFilter(watched QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QObject_PTR() != nil {
		convArg0 = watched.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStateMachine11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10554() {
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
