//  header block begin

// +build !minimal

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

// /usr/include/qt/QtCore/qabstracttransition.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAnimation(QAbstractAnimation *)

/*
Adds the given animation to this transition. The transition does not take ownership of the animation.

See also removeAnimation() and animations().
*/
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

/*
Removes the given animation from this transition.

See also addAnimation().
*/
func (this *QAbstractTransition) RemoveAnimation(animation QAbstractAnimation_ITF /*777 QAbstractAnimation **/) {
	var convArg0 unsafe.Pointer
	if animation != nil && animation.QAbstractAnimation_PTR() != nil {
		convArg0 = animation.QAbstractAnimation_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTransition15removeAnimationEP18QAbstractAnimation", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10300() {
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
