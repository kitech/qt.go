package qtcore

// /usr/include/qt/QtCore/qpauseanimation.h
// #include <qpauseanimation.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QPauseAnimation) InheritEvent(f func(e *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void updateCurrentTime(int)
func (this *QPauseAnimation) InheritUpdateCurrentTime(f func(arg0 int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateCurrentTime", f)
}

type QPauseAnimation struct {
	*QAbstractAnimation
}

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

// /usr/include/qt/QtCore/qpauseanimation.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPauseAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QPauseAnimation10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(QObject *)
func NewQPauseAnimation(parent *QObject /*777 QObject **/) *QPauseAnimation {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:58
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPauseAnimation(int, QObject *)
func NewQPauseAnimation_1(msecs int, parent *QObject /*777 QObject **/) *QPauseAnimation {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", qtrt.FFI_TYPE_POINTER, msecs, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPauseAnimation()
func DeleteQPauseAnimation(this *QPauseAnimation) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qpauseanimation.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int duration()
func (this *QPauseAnimation) Duration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QPauseAnimation8durationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qpauseanimation.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDuration(int)
func (this *QPauseAnimation) SetDuration(msecs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation11setDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:65
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QPauseAnimation) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qpauseanimation.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updateCurrentTime(int)
func (this *QPauseAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QPauseAnimation17updateCurrentTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
