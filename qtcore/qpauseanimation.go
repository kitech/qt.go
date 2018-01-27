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
import "qt.go/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QPauseAnimation) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:0
// Public
// void QPauseAnimation(QObject *)
func NewQPauseAnimation(parent *QObject /*777 QObject **/) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:58
// index:1
// Public
// void QPauseAnimation(int, QObject *)
func NewQPauseAnimation_1(msecs int, parent *QObject /*777 QObject **/) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, msecs, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPauseAnimationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qpauseanimation.h:59
// index:0
// Public virtual
// void ~QPauseAnimation()
func DeleteQPauseAnimation(*QPauseAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:61
// index:0
// Public virtual
// int duration()
func (this *QPauseAnimation) Duration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation8durationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qpauseanimation.h:62
// index:0
// Public
// void setDuration(int)
func (this *QPauseAnimation) SetDuration(msecs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation11setDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:65
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QPauseAnimation) Event(e *QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qpauseanimation.h:66
// index:0
// Protected virtual
// void updateCurrentTime(int)
func (this *QPauseAnimation) UpdateCurrentTime(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation17updateCurrentTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
