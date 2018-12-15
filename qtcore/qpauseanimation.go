// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qpauseanimation.h
// #include <qpauseanimation.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QPauseAnimation) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QPauseAnimation) InheritUpdateCurrentTime(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

/*

 */
type QPauseAnimation struct {
	*QAbstractAnimation
}
type QPauseAnimation_ITF interface {
	QAbstractAnimation_ITF
	QPauseAnimation_PTR() *QPauseAnimation
}

func (ptr *QPauseAnimation) QPauseAnimation_PTR() *QPauseAnimation { return ptr }

func (this *QPauseAnimation) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractAnimation.GetCthis()
	}
}
func (this *QPauseAnimation) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractAnimation = NewQAbstractAnimationFromPointer(cthis)
}
func NewQPauseAnimationFromPointer(cthis unsafe.Pointer) *QPauseAnimation {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QPauseAnimation{bcthis0}
}
func (*QPauseAnimation) NewFromPointer(cthis unsafe.Pointer) *QPauseAnimation {
	return NewQPauseAnimationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qpauseanimation.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QPauseAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QPauseAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpauseanimation.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(QObject *)

/*
Constructs a QPauseAnimation. parent is passed to QObject's constructor. The default duration is 0.
*/
func (*QPauseAnimation) NewForInherit(parent QObject_ITF /*777 QObject **/) *QPauseAnimation {
	return NewQPauseAnimation(parent)
}
func NewQPauseAnimation(parent QObject_ITF /*777 QObject **/) *QPauseAnimation {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPauseAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(QObject *)

/*
Constructs a QPauseAnimation. parent is passed to QObject's constructor. The default duration is 0.
*/
func (*QPauseAnimation) NewForInheritp() *QPauseAnimation {
	return NewQPauseAnimationp()
}
func NewQPauseAnimationp() *QPauseAnimation {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPauseAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(int, QObject *)

/*
Constructs a QPauseAnimation. parent is passed to QObject's constructor. The default duration is 0.
*/
func (*QPauseAnimation) NewForInherit1(msecs int, parent QObject_ITF /*777 QObject **/) *QPauseAnimation {
	return NewQPauseAnimation1(msecs, parent)
}
func NewQPauseAnimation1(msecs int, parent QObject_ITF /*777 QObject **/) *QPauseAnimation {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", qtrt.FFI_TYPE_POINTER, msecs, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPauseAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(int, QObject *)

/*
Constructs a QPauseAnimation. parent is passed to QObject's constructor. The default duration is 0.
*/
func (*QPauseAnimation) NewForInherit1p(msecs int) *QPauseAnimation {
	return NewQPauseAnimation1p(msecs)
}
func NewQPauseAnimation1p(msecs int) *QPauseAnimation {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", qtrt.FFI_TYPE_POINTER, msecs, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QPauseAnimation")
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPauseAnimation()

/*

 */
func DeleteQPauseAnimation(this *QPauseAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qpauseanimation.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration() const

/*

 */
func (this *QPauseAnimation) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QPauseAnimation8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpauseanimation.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuration(int)

/*

 */
func (this *QPauseAnimation) SetDuration(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation11setDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:64
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QAbstractAnimation::event().
*/
func (this *QPauseAnimation) Event(e QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpauseanimation.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)

/*
Reimplemented from QAbstractAnimation::updateCurrentTime().
*/
func (this *QPauseAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
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
}

//  keep block end
