//  header block begin
// /usr/include/qt/QtCore/qpauseanimation.h
// #include <qpauseanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QPauseAnimation struct {
	*QAbstractAnimation
}

func (this *QPauseAnimation) GetCthis() unsafe.Pointer {
	return this.QAbstractAnimation.GetCthis()
}

// /usr/include/qt/QtCore/qpauseanimation.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPauseAnimation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:0
// void QPauseAnimation(class QObject *)
func NewQPauseAnimation(parent unsafe.Pointer) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(cthis)
	return gothis
}
func NewQPauseAnimationFromPointer(cthis unsafe.Pointer) *QPauseAnimation {
	bcthis0 := NewQAbstractAnimationFromPointer(cthis)
	return &QPauseAnimation{bcthis0}
}

// /usr/include/qt/QtCore/qpauseanimation.h:58
// index:1
// void QPauseAnimation(int, class QObject *)
func NewQPauseAnimation_1(msecs int, parent unsafe.Pointer) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &msecs, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:59
// index:0
// virtual
// void ~QPauseAnimation()
func DeleteQPauseAnimation(*QPauseAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:61
// index:0
// virtual
// int duration()
func (this *QPauseAnimation) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation8durationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:62
// index:0
// void setDuration(int)
func (this *QPauseAnimation) SetDuration(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation11setDurationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:65
// index:0
// virtual
// bool event(class QEvent *)
func (this *QPauseAnimation) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:66
// index:0
// virtual
// void updateCurrentTime(int)
func (this *QPauseAnimation) UpdateCurrentTime(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
